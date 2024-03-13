package service

import (
	"github.com/pkg/errors"
	"os"

	"github.com/sirupsen/logrus"

	"deuna.com/payment/bank/api"
	"deuna.com/payment/bank/models/interfaces"
)

var bankService Banker

func DefaultBankService() Banker {
	if bankService == nil {
		bankService = NewBankService()
	}

	return bankService
}

type Banker interface {
	GetAccount(bankName, accountNumber string) (interfaces.Account, error)
}

type Bank struct {
	api *api.BankService
}

func NewBankService() Banker {
	bankServerURL := os.Getenv("BANK_SERVER_URL")
	if bankServerURL == "" {
		logrus.Infof("BANK_SERVER_URL not found, using default localhost:8080")

		bankServerURL = "http://localhost:8080"
	}

	return &Bank{
		api: api.NewBankService(bankServerURL),
	}
}

func (b *Bank) GetAccount(bankName, accountNumber string) (interfaces.Account, error) {
	account, err := b.api.GetAccount(bankName, accountNumber)
	if err != nil {
		return nil, errors.Wrap(err, "service.bank.get_account: getting account")
	}

	return account, nil
}
