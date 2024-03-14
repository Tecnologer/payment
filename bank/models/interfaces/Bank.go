package interfaces

type Bank interface {
	GetName() string
	GetAccount(accountNumber string) (Account, error)
	Transfer(originAccountNumber, destinationAccountNumber string, amount float32) error
}
