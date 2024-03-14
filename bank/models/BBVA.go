package models

import (
	berrors "deuna.com/payment/bank/errors"
	"deuna.com/payment/bank/models/interfaces"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var staticBBVA *BBVA

func init() {
	staticBBVA = &BBVA{
		Name: "BBVA",
		Accounts: map[string]*Account{
			"123456": {
				Client: &Client{
					Name: "John Nommensen",
				},
				Number:  "123456",
				Balance: 1000,
			},
			"654321": {
				Client: &Client{
					Name: "John Doe",
				},
				Number:  "654321",
				Balance: 10000,
			},
			"111111": {
				Client: &Client{
					Name: "Nexus Innovate",
				},
				Number:  "111111",
				Balance: 20000,
			},
			"222222": {
				Client: &Client{
					Name: "Deuna",
				},
				Number:  "222222",
				Balance: 8980,
			},
		},
	}
}

type BBVA struct {
	Name     string              `json:"name"`
	Accounts map[string]*Account `json:"-"`
}

func NewBBVA() *BBVA {
	return staticBBVA
}

func (b *BBVA) GetName() string {
	return b.Name
}

func (b *BBVA) GetAccount(accountNumber string) (interfaces.Account, error) {
	account, ok := b.Accounts[accountNumber]
	if !ok {
		logrus.Errorf("BBVA: Account %s not found", accountNumber)
		return nil, berrors.AccountNotFound
	}

	account.BankName = b.Name

	return account, nil
}

func (b *BBVA) Transfer(originAccountNumber, destinationAccountNumber string, amount float32) error {
	// update references to the accounts
	origin, err := b.GetAccount(originAccountNumber)
	if err != nil {
		return errors.Errorf("BBVA: Origin account %s not found", originAccountNumber)
	}

	destination, err := b.GetAccount(destinationAccountNumber)
	if err != nil {
		return errors.Errorf("BBVA: Destination account %s not found", destinationAccountNumber)
	}

	if err := origin.Withdraw(amount); err != nil {
		return err
	}

	_ = destination.Deposit(amount)

	logrus.Infof("BBVA: Transfer from %s to %s for %f", origin.GetID(), destination.GetID(), amount)

	return nil
}
