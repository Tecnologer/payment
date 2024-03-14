package models

import (
	"deuna.com/payment/gatepay/src/validator"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	*gorm.Model
	Name          string    `json:"name"                  validate:"required" gorm:"not null"`
	BankName      string    `json:"bank_name"             validate:"required" gorm:"not null"`
	AccountNumber string    `json:"account_number"        validate:"required" gorm:"not null"`
	OwnerName     string    `json:"owner_name"`
	OwnerEmail    string    `json:"owner_email"           validate:"required" gorm:"not null"`
	CustomerID    *uint     `json:"customer_id,omitempty"`
	Customer      *Customer `json:"customer,omitempty"                        gorm:"foreignKey:CustomerID;"`
	MerchantID    *uint     `json:"merchant_id,omitempty"`
	Merchant      *Merchant `json:"merchant,omitempty"                        gorm:"foreignKey:MerchantID;"`
}

func (pm *PaymentMethod) Validate() error {
	return validator.ValidateObject(pm)
}
