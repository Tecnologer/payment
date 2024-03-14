package business

import (
	"context"

	"deuna.com/payment/gatepay/src/business/dao"
	"deuna.com/payment/gatepay/src/models"
	"deuna.com/payment/gatepay/src/service"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	*Business
	daoPaymentMethod *dao.PaymentMethod
	daoCustomer      *dao.Customer
	daoMerchantUser  *dao.MerchantUser
	bank             service.Banker
}

func NewPaymentMethod(db *gorm.DB, ctx context.Context) *PaymentMethod {
	return &PaymentMethod{
		Business:         NewBusiness(ctx),
		daoPaymentMethod: dao.NewPaymentMethod(db),
		daoCustomer:      dao.NewCustomer(db),
		daoMerchantUser:  dao.NewMerchantUser(db),
		bank:             service.NewBankService(),
	}
}

func (p *PaymentMethod) Create(inputPaymentMethod *models.PaymentMethod) (*models.PaymentMethod, error) {
	err := p.setOwner(inputPaymentMethod)
	if err != nil {
		return nil, errors.Wrap(err, "business.paymentMethod.create: setting owner")
	}

	if !p.existsAccountInBank(inputPaymentMethod) {
		return nil, errors.Errorf(
			"business.paymentMethod.create: the user with email %s does not have an account with number %s in bank %s",
			inputPaymentMethod.OwnerEmail,
			inputPaymentMethod.AccountNumber,
			inputPaymentMethod.BankName,
		)
	}

	paymentMethod, err := p.daoPaymentMethod.InsertIfNoExists(inputPaymentMethod)
	if err != nil {
		return nil, err
	}

	return paymentMethod, nil
}

func (p *PaymentMethod) existsAccountInBank(paymentMethod *models.PaymentMethod) bool {
	_, err := p.bank.GetAccount(
		p.Context,
		paymentMethod.OwnerName,
		paymentMethod.BankName,
		paymentMethod.AccountNumber,
	)
	if err != nil {
		logrus.WithError(err).Error("business.paymentMethod.exists_account_in_bank: validating account")
	}

	return err == nil
}

func (p *PaymentMethod) setOwner(paymentMethod *models.PaymentMethod) error {
	customer, err := p.daoCustomer.FindByEmail(paymentMethod.OwnerEmail)
	if err != nil {
		return errors.Wrap(err, "business.paymentMethod.set_owner: getting customer by email")
	}

	if customer != nil {
		paymentMethod.CustomerID = &customer.ID
		paymentMethod.OwnerName = customer.Name

		return nil
	}

	merchantUser, err := p.daoMerchantUser.FindByEmail(paymentMethod.OwnerEmail)
	if err != nil {
		return errors.Wrap(err, "business.paymentMethod.set_owner: getting merchant by email")
	}

	if merchantUser != nil {
		if merchantUser.Role == models.Staff {
			return errors.Errorf(
				"business.paymentMethod.set_owner: the user with email %s is not allowed to create payment methods for merchant %s",
				paymentMethod.OwnerEmail,
				merchantUser.Merchant.Name,
			)
		}

		paymentMethod.MerchantID = &merchantUser.MerchantID
		paymentMethod.OwnerName = merchantUser.Merchant.Name

		return nil
	}

	return errors.Errorf(
		"business.paymentMethod.set_owner: the email %s does not belong to a customer or merchant",
		paymentMethod.OwnerEmail,
	)
}
