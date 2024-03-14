package dao

import (
	"deuna.com/payment/gatepay/src/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Item struct {
	db *gorm.DB
}

func NewItem(db *gorm.DB) *Item {
	return &Item{
		db: db,
	}
}

func (i *Item) InsertIfNotExists(merchant *models.Merchant, item *models.Item) (*models.Item, error) {
	tx := i.db.Where("description = ? AND merchant_id = ?", item.Description, merchant.ID).First(&item)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(tx.Error, "dao.item.insert_if_not_exists: finding item")
	}

	item.MerchantID = merchant.ID

	tx = i.db.Save(item)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "dao.item.insert_if_not_exists: inserting item")
	}

	return item, nil
}
