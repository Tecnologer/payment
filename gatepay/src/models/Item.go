package models

import (
	"deuna.com/payment/gatepay/src/validator"
	"gorm.io/gorm"
)

type Item struct {
	*gorm.Model
	MerchantID  uint      `json:"merchant_id" validate:"gt=0"     gorm:"not null"`
	Description string    `json:"description" validate:"required" gorm:"not null"`
	Price       float32   `json:"price"       validate:"gt=0"     gorm:"not null"`
	Quantity    float32   `json:"quantity"                        gorm:"not null"`
	Merchant    *Merchant `json:"merchant"                        gorm:"foreignKey:MerchantID"`
}

func (i *Item) Validate() error {
	return validator.ValidateObject(i)
}
