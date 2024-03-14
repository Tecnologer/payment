package interfaces

type Account interface {
	GetID() string
	GetBankName() string
	Withdraw(amount float32) error
	Deposit(amount float32) error
	GetOwnerName() string
}
