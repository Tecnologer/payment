package dao

import (
	"deuna.com/payment/gatepay/src/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Payment struct {
	db *gorm.DB
}

func NewPayment(db *gorm.DB) *Payment {
	return &Payment{db: db}
}

func (p *Payment) Insert(inputPayment *models.Payment) (*models.Payment, error) {
	tx := p.db.Save(inputPayment)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "dao.payment.insert: inserting payment")
	}

	inputPayment, err := p.FindByID(inputPayment.ID)
	if err != nil {
		return nil, errors.Wrap(err, "dao.payment.insert: retrieving saved payment")
	}

	return inputPayment, nil
}

func (p *Payment) FindByID(id uint) (*models.Payment, error) {
	var payment *models.Payment

	tx := p.db.
		Preload("OriginPaymentMethod").
		Preload("Customer").
		Preload("Merchant").
		Preload("DestinationPaymentMethod").
		Preload("Items").
		First(&payment, id)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "dao.payment.find_by_id")
	}

	return payment, nil
}
