package service

import "deuna.com/payment/bank/models/interfaces"

type BankServiceMock byte

func SetBankServiceMockAsDefault() {
	bankService = NewBankMock()
}

func NewBankMock() Banker {
	return new(BankServiceMock)
}

type BankMockModel struct {
	Name string
}

func (b *BankMockModel) GetName() string {
	return b.Name
}
func (b *BankMockModel) GetAccount(accountNumber string) (interfaces.Account, error) {
	return &Account{
		BankName: b.Name,
		Number:   accountNumber,
	}, nil
}

func (b *BankMockModel) Payment(origin, destination interfaces.Account, amount float32) error {
	return nil
}

type Account struct {
	BankName string
	Number   string
}

func (b *Account) GetID() string {
	return "mock-account"
}

func (b *Account) GetBank() interfaces.Bank {
	return &BankMockModel{
		Name: b.BankName,
	}
}

func (b *Account) Withdraw(amount float32) error {
	return nil
}

func (b *Account) Deposit(amount float32) error {
	return nil
}

func (b *BankServiceMock) GetAccount(bankName, accountNumber string) (interfaces.Account, error) {
	return &Account{
		BankName: bankName,
		Number:   accountNumber,
	}, nil
}
