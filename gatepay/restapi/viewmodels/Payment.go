package viewmodels

import (
	"deuna.com/payment/gatepay/src/models"
)

type Payment struct {
	OriginPaymentMethodID      uint           `json:"origin_payment_method_id"`
	DestinationPaymentMethodID uint           `json:"destination_payment_method_id"`
	Amount                     float32        `json:"amount"`
	Items                      []*models.Item `json:"items"`
}

func (p *Payment) ParseToModelPayment() *models.Payment {
	payment := &models.Payment{
		OriginPaymentMethodID:      p.OriginPaymentMethodID,
		DestinationPaymentMethodID: p.DestinationPaymentMethodID,
		Amount:                     p.Amount,
		Items:                      make([]*models.PaymentItem, len(p.Items)),
	}

	for i, item := range p.Items {
		payment.Items[i] = &models.PaymentItem{
			Item: item,
		}
	}

	return payment
}
