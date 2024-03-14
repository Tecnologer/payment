package business

import (
	"context"

	"deuna.com/payment/bank/models/interfaces"
	"deuna.com/payment/gatepay/src/business/dao"
	"deuna.com/payment/gatepay/src/models"
	"deuna.com/payment/gatepay/src/service"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Payment struct {
	*Business
	dao              *dao.Payment
	daoPaymentMethod *dao.PaymentMethod
	bank             service.Banker
	daoCustomer      *dao.Customer
	daoMerchant      *dao.Merchant
}

type accounts struct {
	origin      interfaces.Account
	destination interfaces.Account
}

func NewPayment(db *gorm.DB, ctx context.Context) *Payment {
	return &Payment{
		Business:         NewBusiness(ctx),
		dao:              dao.NewPayment(db),
		bank:             service.DefaultBankService(),
		daoPaymentMethod: dao.NewPaymentMethod(db),
		daoCustomer:      dao.NewCustomer(db),
		daoMerchant:      dao.NewMerchant(db),
	}
}

func (p *Payment) Register(customerEmail string, inputPayment *models.Payment) (*models.Payment, error) {
	customer, err := p.daoCustomer.FindByEmail(customerEmail)
	if err != nil {
		return nil, errors.Wrap(err, "business.payment.register: getting customer")
	}

	if customer == nil {
		return nil, errors.New("business.payment.register: customer not found")
	}

	err = p.retrievePaymentMethods(inputPayment)
	if err != nil {
		return nil, errors.Wrap(err, "business.payment.register: retrieving payment methods")
	}

	accounts, err := p.paymentBankAccounts(inputPayment)
	if err != nil {
		return nil, errors.Wrap(err, "business.payment.register: getting accounts")
	}

	if customer.Name != accounts.origin.GetOwnerName() {
		return nil, errors.New("business.payment.register: the account %s does not belong to the customer %s")
	}

	newPayment, err := p.dao.Insert(inputPayment)
	if err != nil {
		return nil, errors.Wrap(err, "business.payment.register: inserting payment")
	}

	if err := p.bank.Transfer(p.Context, accounts.origin, accounts.destination, inputPayment.Amount); err != nil {
		return nil, errors.Wrap(err, "business.payment.register: transferring money")
	}

	return newPayment, nil
}

func (p *Payment) retrievePaymentMethods(inputPayment *models.Payment) error {
	originPaymentMethod, err := p.daoPaymentMethod.FindByID(inputPayment.OriginPaymentMethodID)
	if err != nil {
		return errors.Wrap(err, "business.payment.retrieve_payment_methods: retrieving origin payment method")
	}

	if originPaymentMethod == nil {
		return errors.New("business.payment.retrieve_payment_methods: origin payment method not found")
	}

	inputPayment.OriginPaymentMethod = originPaymentMethod

	destinationPaymentMethod, err := p.daoPaymentMethod.FindByID(inputPayment.DestinationPaymentMethodID)
	if err != nil {
		return errors.Wrap(
			err,
			"business.payment.retrieve_payment_methods: retrieving destination payment method",
		)
	}

	if destinationPaymentMethod == nil {
		return errors.New("business.payment.retrieve_payment_methods: destination payment method not found")
	}

	inputPayment.DestinationPaymentMethod = destinationPaymentMethod

	return nil
}

func (p *Payment) paymentBankAccounts(inputPayment *models.Payment) (a accounts, err error) {
	a.origin, err = p.bank.GetAccount(
		p.Context,
		inputPayment.OriginPaymentMethod.OwnerName,
		inputPayment.OriginPaymentMethod.BankName,
		inputPayment.OriginPaymentMethod.AccountNumber,
	)
	if err != nil {
		return a, errors.Wrap(err, "business.payment.payment_accounts: getting origin account")
	}

	a.destination, err = p.bank.GetAccount(
		p.Context,
		inputPayment.DestinationPaymentMethod.OwnerName,
		inputPayment.DestinationPaymentMethod.BankName,
		inputPayment.DestinationPaymentMethod.AccountNumber,
	)
	if err != nil {
		return a, errors.Wrap(err, "business.payment.payment_accounts: getting destination account")
	}

	return
}

func (p *Payment) PaymentsByEmailOwner(customerEmail string) ([]*models.Payment, error) {
	isCustomer, err := p.daoCustomer.IsEmailCustomer(customerEmail)
	if err != nil {
		return nil, errors.Wrap(err, "business.payment.payments_by_email_owner: checking if is customer")
	}

	if isCustomer {
		return p.dao.PaymentsSentByEmailOwner(customerEmail)
	}

	return p.dao.PaymentsReceivedByEmailUser(customerEmail)
}

func (p *Payment) Refund(userEmail string, paymentID uint) error {
	payment, err := p.ValidatePaymentID(paymentID)
	if err != nil {
		return errors.Wrap(err, "business.payment.refund: validating payment id")
	}

	err = p.validateUserEmailForRefund(userEmail, payment)
	if err != nil {
		return errors.Wrap(err, "business.payment.refund: validating user email")
	}

	accounts, err := p.paymentBankAccounts(payment)
	if err != nil {
		return errors.Wrap(err, "business.payment.refund: getting accounts")
	}

	err = p.dao.Refund(paymentID)
	if err != nil {
		return errors.Wrap(err, "business.payment.refund: refunding payment")
	}

	if err := p.bank.Transfer(p.Context, accounts.destination, accounts.origin, payment.Amount); err != nil {
		return errors.Wrap(err, "business.payment.refund: transferring money")
	}

	return nil
}

func (p *Payment) ValidatePaymentID(paymentID uint) (*models.Payment, error) {
	payment, err := p.dao.FindByID(paymentID)
	if err != nil {
		return nil, errors.Wrap(err, "business.payment.refund: finding payment")
	}

	if payment == nil {
		return nil, errors.New("business.payment.refund: payment not found")
	}

	if payment.Status != models.PaymentStatusApproved {
		return nil, errors.New("business.payment.refund: payment is already refunded")
	}

	return payment, nil
}

func (p *Payment) validateUserEmailForRefund(userEmail string, payment *models.Payment) error {
	isCustomer, err := p.daoCustomer.IsEmailCustomer(userEmail)
	if err != nil {
		return errors.Wrap(err, "business.payment.refund: checking if is customer")
	}

	// both customer and merchant can refund a payment
	if isCustomer {
		if payment.OriginPaymentMethod.Customer.Email != userEmail {
			return errors.New("business.payment.refund: the payment does not belong to the customer")
		}
	} else {
		if !payment.DestinationPaymentMethod.Merchant.EmailBelongsToMerchant(userEmail) {
			return errors.New("business.payment.refund: the payment does not belong to the merchant")
		}
	}

	return nil
}
