package interfaces

type Account interface {
	GetID() string
	GetBank() Bank
	Withdraw(amount float32) error
	Deposit(amount float32) error
}
