package viewmodels

import "github.com/pkg/errors"

type Account struct {
	BankName  string `json:"bank_name"`
	Number    string `json:"number"`
	OwnerName string `json:"owner_name"`
}

func (a *Account) Validate() error {
	if a.BankName == "" {
		return errors.New("bank name is required")
	}

	if a.Number == "" {
		return errors.New("number is required")
	}

	if a.OwnerName == "" {
		return errors.New("owner name is required")
	}

	return nil
}
