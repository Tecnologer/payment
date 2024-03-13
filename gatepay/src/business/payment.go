package business

import (
	"deuna.com/payment/bank/models/interfaces"
	"deuna.com/payment/gatepay/src/business/dao"
	"deuna.com/payment/gatepay/src/models"
	"deuna.com/payment/gatepay/src/service"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Payment struct {
	dao  *dao.Payment
	bank service.Banker
}

type accounts struct {
	origin      interfaces.Account
	destination interfaces.Account
}

func NewPayment(db *gorm.DB) *Payment {
	return &Payment{
		dao:  dao.NewPayment(db),
		bank: service.DefaultBankService(),
	}
}

func (p *Payment) Register(inputPayment *models.Payment) (*models.Payment, error) {
	//accounts, err := p.paymentAccounts(inputPayment)
	//if err != nil {
	//	return nil, errors.Wrap(err, "business.payment.register: getting accounts")
	//}

	return p.dao.Insert(inputPayment)
}

func (p *Payment) paymentAccounts(inputPayment *models.Payment) (a accounts, err error) {
	a.origin, err = p.bank.GetAccount(
		inputPayment.OriginPaymentMethod.BankName,
		inputPayment.OriginPaymentMethod.AccountNumber,
	)
	if err != nil {
		return a, errors.Wrap(err, "business.payment.payment_accounts: getting origin account")
	}

	a.destination, err = p.bank.GetAccount(
		inputPayment.DestinationPaymentMethod.BankName,
		inputPayment.DestinationPaymentMethod.AccountNumber,
	)
	if err != nil {
		return a, errors.Wrap(err, "business.payment.payment_accounts: getting destination account")
	}

	return
}
