package dao

import (
	"deuna.com/payment/gatepay/src/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type PaymentItem struct {
	db   *gorm.DB
	item *Item
}

func NewPaymentItem(db *gorm.DB) *PaymentItem {
	return &PaymentItem{
		db:   db,
		item: NewItem(db),
	}
}

func (p *PaymentItem) Insert(
	paymentID uint,
	merchant *models.Merchant,
	paymentItems ...*models.PaymentItem,
) ([]*models.PaymentItem, error) {
	if len(paymentItems) == 0 {
		return nil, nil
	}

	for _, paymentItem := range paymentItems {
		item, err := p.item.InsertIfNotExists(merchant, paymentItem.Item)
		if err != nil {
			return nil, errors.Wrap(err, "dao.payment_item.insert: inserting item")
		}

		paymentItem.PaymentID = paymentID
		paymentItem.ItemID = item.ID
		paymentItem.Quantity = item.Quantity
		paymentItem.Price = item.Price
		paymentItem.Item = nil

		tx := p.db.Save(&paymentItem)
		if tx.Error != nil {
			return nil, errors.Wrap(tx.Error, "dao.payment_item.insert: inserting payment item")
		}
	}

	return paymentItems, nil
}
