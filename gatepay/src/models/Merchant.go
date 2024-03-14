package models

import "gorm.io/gorm"

type Merchant struct {
	*gorm.Model
	Name  string          `json:"name"  validate:"required"`
	Users []*MerchantUser `json:"users" validate:"required" gorm:"foreignKey:MerchantID"`
	Items []*Item         `json:"items"                     gorm:"foreignKey:MerchantID"`
}

type MerchantUser struct {
	*Person
	Role       MerchantRole `json:"role"`
	MerchantID uint         `json:"merchant_id"`
	Merchant   *Merchant    `json:"merchant"    gorm:"foreignKey:MerchantID"`
}

func (mu *Merchant) EmailBelongsToMerchant(email string) bool {
	for _, user := range mu.Users {
		if user.Email == email {
			return true
		}
	}

	return false
}
