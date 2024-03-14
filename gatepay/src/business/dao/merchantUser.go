package dao

import (
	"deuna.com/payment/gatepay/src/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MerchantUser struct {
	db *gorm.DB
}

func NewMerchantUser(db *gorm.DB) *MerchantUser {
	return &MerchantUser{db: db}
}

func (m *MerchantUser) FindByEmail(email string) (*models.MerchantUser, error) {
	var user *models.MerchantUser

	tx := m.db.
		Preload("Merchant").
		Where("email = ?", email).
		First(&user)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(tx.Error, "dao.merchantUser.find_by_email: finding merchant user by email")
	}

	return user, nil
}
