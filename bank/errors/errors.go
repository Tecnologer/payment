package errors

type Error string

const (
	InsufficientFunds     = Error("insufficient_funds")
	AccountNotFound       = Error("account_not_found")
	InvalidDepositAmount  = Error("invalid_deposit_amount")
	InvalidWithdrawAmount = Error("invalid_withdraw_amount")
)

func (e Error) Error() string {
	return string(e)
}
