package service

import (
	"context"

	"deuna.com/payment/bank/models/interfaces"
)

type BankServiceMock byte

func SetBankServiceMockAsDefault() {
	bankService = NewBankMock()
}

func NewBankMock() Banker {
	return new(BankServiceMock)
}

func (b *BankServiceMock) GetAccount(
	ctx context.Context,
	ownerName, bankName, number string,
) (interfaces.Account, error) {
	return &Account{
		BankName:  bankName,
		Number:    number,
		OwnerName: ownerName,
	}, nil
}

func (b *BankServiceMock) Transfer(ctx context.Context, origin, destination interfaces.Account, amount float32) error {
	// TODO: implement transfer
	return nil
}

type Account struct {
	BankName  string
	Number    string
	OwnerName string
}

func (b *Account) GetID() string {
	return "mock-account"
}

func (b *Account) GetBankName() string {
	return b.BankName
}

func (b *Account) Withdraw(amount float32) error {
	return nil
}

func (b *Account) Deposit(amount float32) error {
	return nil
}

func (b *Account) GetOwnerName() string {
	return b.OwnerName
}
