package models

import (
	"deuna.com/payment/gatepay/src/validator"
	"gorm.io/gorm"
)

type Payment struct {
	*gorm.Model
	CustomerID                 uint           `json:"customer_id"                   validate:"gte=0" gorm:"not null"`
	MerchantID                 uint           `json:"merchant_id"                   validate:"gte=0" gorm:"not null"`
	OriginPaymentMethodID      uint           `json:"origin_payment_method_id"      validate:"gte=0" gorm:"not null"`
	DestinationPaymentMethodID uint           `json:"destination_payment_method_id" validate:"gte=0" gorm:"not null"`
	Amount                     float32        `json:"amount"                        validate:"gte=0" gorm:"not null"`
	OriginPaymentMethod        *PaymentMethod `json:"origin_payment_method"                          gorm:"foreignKey:OriginPaymentMethodID;"`
	Customer                   *Customer      `json:"customer"                                       gorm:"foreignKey:CustomerID;"`
	Merchant                   *Merchant      `json:"merchant"                                       gorm:"foreignKey:MerchantID;"`
	DestinationPaymentMethod   *PaymentMethod `json:"destination_payment_method"                     gorm:"foreignKey:DestinationPaymentMethodID;"`
	Items                      []*PaymentItem `json:"items"                                          gorm:"foreignKey:PaymentID;"`
}

func (p *Payment) Validate() error {
	return validator.ValidateObject(p)
}
