package service

import (
	"context"
	"os"

	"github.com/pkg/errors"

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
	GetAccount(ctx context.Context, ownerName, bankName, number string) (interfaces.Account, error)
	Transfer(ctx context.Context, origin, destination interfaces.Account, amount float32) error
}

type Bank struct {
	api *api.BankService
}

func NewBankService() Banker {
	bankServerURL := os.Getenv("BANK_SERVER_URL")
	if bankServerURL == "" {
		logrus.Warning("BANK_SERVER_URL not found, using default localhost:8080")

		bankServerURL = "http://localhost:8080"
	}

	return &Bank{
		api: api.NewBankService(bankServerURL),
	}
}

func (b *Bank) GetAccount(ctx context.Context, ownerName, bankName, number string) (interfaces.Account, error) {
	request := api.NewAccountRequest(ctx, ownerName, bankName, number)

	account, err := b.api.GetAccount(request)
	if err != nil {
		return nil, errors.Wrap(err, "service.bank.get_account: getting account")
	}

	return account, nil
}

func (b *Bank) Transfer(ctx context.Context, origin, destination interfaces.Account, amount float32) error {
	request := api.NewTransferRequest(ctx, origin, destination, amount)

	err := b.api.Transfer(request)
	if err != nil {
		return errors.Wrap(err, "service.bank.transfer: transferring money")
	}

	return nil
}
