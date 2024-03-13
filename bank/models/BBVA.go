package models

import (
	"deuna.com/payment/bank/errors"
	"deuna.com/payment/bank/models/interfaces"
	"github.com/sirupsen/logrus"
)

type BBVA struct {
	Name     string              `json:"name"`
	Accounts map[string]*Account `json:"-"`
}

func NewBBVA() *BBVA {
	return &BBVA{
		Name: "BBVA",
		Accounts: map[string]*Account{
			"123456": {
				Client: &Client{
					Name: "Tecnologer",
				},
				Number:  "123456",
				Balance: 1000,
			},
			"654321": {
				Client: &Client{
					Name: "Deuna",
				},
				Number:  "654321",
				Balance: 10000,
			},
		},
	}
}

func (b *BBVA) GetName() string {
	return b.Name
}

func (b *BBVA) GetAccount(accountNumber string) (interfaces.Account, error) {
	account, ok := b.Accounts[accountNumber]
	if !ok {
		return nil, errors.AccountNotFound
	}

	return account, nil
}

func (b *BBVA) Payment(origin, destination interfaces.Account, amount float32) error {
	if err := origin.Withdraw(amount); err != nil {
		return err
	}

	_ = destination.Deposit(amount)

	logrus.Infof("Payment from %s to %s for %f", origin.GetID(), destination.GetID(), amount)

	return nil
}
