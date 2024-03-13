package models

import (
	"deuna.com/payment/gatepay/src/validator"
	"gorm.io/gorm"
)

type PaymentItem struct {
	*gorm.Model
	PaymentID uint     `json:"payment_id" validate:"gte=0" gorm:"not null"`
	ItemID    uint     `json:"item_id"    validate:"gte=0" gorm:"not null"`
	Quantity  float32  `json:"quantity"   validate:"gte=0" gorm:"not null"`
	Price     float32  `json:"price"      validate:"gte=0" gorm:"not null"`
	Payment   *Payment `json:"payment"                     gorm:"foreignKey:PaymentID;"`
	Item      *Item    `json:"item"                        gorm:"foreignKey:ItemID;"`
}

func (pi *PaymentItem) Validate() error {
	return validator.ValidateObject(pi)
}
