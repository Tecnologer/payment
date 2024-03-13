package models

import (
	"deuna.com/payment/bank/errors"
	"deuna.com/payment/bank/models/interfaces"
)

type Client struct {
	Name string `json:"name"`
}

type Account struct {
	Client   *Client         `json:"client"`
	BankName interfaces.Bank `json:"bank_name"`
	Number   string          `json:"number"`
	Balance  float32         `json:"balance"`
}

func (a *Account) GetID() string {
	return a.Number
}

func (a *Account) GetBank() interfaces.Bank {
	return a.BankName
}

func (a *Account) Withdraw(amount float32) error {
	if amount < 0 {
		return errors.InvalidWithdrawAmount
	}

	if a.Balance < amount {
		return errors.InsufficientFunds
	}

	a.Balance -= amount

	return nil
}

func (a *Account) Deposit(amount float32) error {
	if amount < 0 {
		return errors.InvalidDepositAmount
	}

	a.Balance += amount

	return nil
}
