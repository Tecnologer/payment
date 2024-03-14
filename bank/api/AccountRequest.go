package api

import "context"

type AccountRequest struct {
	Request
	OwnerName string `json:"owner_name"`
	BankName  string `json:"bank_name"`
	Number    string `json:"number"`
}

func NewAccountRequest(ctx context.Context, ownerName, bankName, number string) *AccountRequest {
	return &AccountRequest{
		Request:   Request{Context: ctx},
		OwnerName: ownerName,
		BankName:  bankName,
		Number:    number,
	}
}
