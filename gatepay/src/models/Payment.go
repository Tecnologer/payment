package models

import (
	"deuna.com/payment/gatepay/src/validator"
	"gorm.io/gorm"
)

type Payment struct {
	*gorm.Model
	OriginPaymentMethodID      uint           `json:"origin_payment_method_id"      validate:"gt=0" gorm:"not null"`
	DestinationPaymentMethodID uint           `json:"destination_payment_method_id" validate:"gt=0" gorm:"not null"`
	Amount                     float32        `json:"amount"                        validate:"gt=0" gorm:"not null"`
	OriginPaymentMethod        *PaymentMethod `json:"origin_payment_method"                         gorm:"foreignKey:OriginPaymentMethodID;"`
	DestinationPaymentMethod   *PaymentMethod `json:"destination_payment_method"                    gorm:"foreignKey:DestinationPaymentMethodID;"`
	Items                      []*PaymentItem `json:"items"                                         gorm:"foreignKey:PaymentID;"`
	Status                     PaymentStatus  `json:"status"                                        gorm:"not null;"`
}

func (p *Payment) Validate() error {
	return validator.ValidateObject(p)
}
