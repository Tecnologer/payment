package interfaces

type Bank interface {
	GetName() string
	GetAccount(accountNumber string) (Account, error)
	Payment(origin, destination Account, amount float32) error
}
