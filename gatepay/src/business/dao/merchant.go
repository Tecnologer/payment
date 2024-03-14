package dao

import (
	"deuna.com/payment/gatepay/src/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Merchant struct {
	db *gorm.DB
}

func NewMerchant(db *gorm.DB) *Merchant {
	return &Merchant{db: db}
}

func (m *Merchant) FindByID(id uint) (*models.Merchant, error) {
	var merchant *models.Merchant

	tx := m.db.First(&merchant, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(tx.Error, "dao.merchant.find_by_id")
	}

	return merchant, nil
}
