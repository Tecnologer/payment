package factory

import (
	"strings"

	"deuna.com/payment/bank/models"
	"deuna.com/payment/bank/models/interfaces"
)

func NewBank(name string) interfaces.Bank {
	name = strings.ToLower(name)

	switch name {
	case "bbva":
		return models.NewBBVA()
	default:
		return nil
	}
}
