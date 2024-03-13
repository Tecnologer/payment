package viewmodels

import (
	"deuna.com/payment/gatepay/src/models"
)

type Payment struct {
	Amount             float32               `json:"amount"`
	Items              []*models.PaymentItem `json:"items"`
	OriginAccount      *Account              `json:"origin_account"`
	DestinationAccount *Account              `json:"destination_account"`
}

type Account struct {
	Number   string `json:"number"`
	BankName string `json:"bank_name"`
}

func (p *Payment) ToPaymentModel() *models.Payment {
	return &models.Payment{
		Amount: p.Amount,
		OriginPaymentMethod: &models.PaymentMethod{
			AccountNumber: p.OriginAccount.Number,
			BankName:      p.OriginAccount.BankName,
		},
		DestinationPaymentMethod: &models.PaymentMethod{
			AccountNumber: p.DestinationAccount.Number,
			BankName:      p.DestinationAccount.BankName,
		},
		Items: p.Items,
	}
}
