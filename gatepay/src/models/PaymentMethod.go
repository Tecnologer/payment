package models

import (
	"deuna.com/payment/gatepay/src/validator"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	*gorm.Model
	Name          string    `json:"name"           validate:"required" gorm:"not null"`
	BankName      string    `json:"bank_name"      validate:"required" gorm:"not null"`
	AccountNumber string    `json:"account_number" validate:"required" gorm:"not null"`
	OwnerName     string    `json:"owner_name"     validate:"required" gorm:"not null"`
	CustomerID    *uint     `json:"customer_id"`
	Customer      *Customer `json:"customer"                           gorm:"foreignKey:CustomerID;"`
	MerchantID    *uint     `json:"merchant_id"`
	Merchant      *Merchant `json:"merchant"                           gorm:"foreignKey:MerchantID;"`
}

func (pm *PaymentMethod) Validate() error {
	return validator.ValidateObject(pm)
}
