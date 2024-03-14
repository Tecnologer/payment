package dao

import (
	"deuna.com/payment/gatepay/src/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Payment struct {
	db          *gorm.DB
	paymentItem *PaymentItem
}

func NewPayment(db *gorm.DB) *Payment {
	return &Payment{
		db:          db,
		paymentItem: NewPaymentItem(db),
	}
}

func (p *Payment) Insert(inputPayment *models.Payment) (*models.Payment, error) {
	var (
		merchant = inputPayment.DestinationPaymentMethod.Merchant
		items    = append([]*models.PaymentItem{}, inputPayment.Items...)
	)

	inputPayment.Items = nil
	inputPayment.DestinationPaymentMethod = nil

	tx := p.db.Save(inputPayment)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "dao.payment.insert: inserting payment")
	}

	_, err := p.paymentItem.Insert(inputPayment.ID, merchant, items...)
	if err != nil {
		return nil, errors.Wrap(err, "dao.payment.insert: inserting payment items")
	}

	inputPayment, err = p.FindByID(inputPayment.ID)
	if err != nil {
		return nil, errors.Wrap(err, "dao.payment.insert: retrieving saved payment")
	}

	return inputPayment, nil
}

func (p *Payment) FindByID(id uint) (*models.Payment, error) {
	var payment *models.Payment

	tx := p.db.
		Preload("OriginPaymentMethod.Customer").
		Preload("DestinationPaymentMethod.Merchant").
		Preload("Items").
		First(&payment, id)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "dao.payment.find_by_id")
	}

	return payment, nil
}

func (p *Payment) PaymentsReceivedByEmailUser(email string) ([]*models.Payment, error) {
	var payments []*models.Payment

	tx := p.db.
		Preload("OriginPaymentMethod.Customer").
		Preload("DestinationPaymentMethod.Merchant").
		Preload("Items").
		Joins("JOIN payment_methods ON payment_methods.id = payments.destination_payment_method_id").
		Joins("JOIN merchants ON merchants.id = payment_methods.merchant_id").
		Joins("JOIN merchant_users ON merchant_users.merchant_id = merchants.id").
		Where("merchant_users.email = ?", email).
		Find(&payments)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "dao.payment.get_payments_by_customer_email")
	}

	return payments, nil
}

func (p *Payment) PaymentsSentByEmailOwner(email string) ([]*models.Payment, error) {
	var payments []*models.Payment

	tx := p.db.
		Preload("OriginPaymentMethod.Customer").
		Preload("DestinationPaymentMethod.Merchant").
		Preload("Items").
		Joins("JOIN payment_methods ON payment_methods.id = payments.origin_payment_method_id").
		Joins("JOIN customers ON customers.id = payment_methods.customer_id").
		Where("customers.email = ?", email).
		Find(&payments)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "dao.payment.get_payments_by_customer_email")
	}

	return payments, nil
}

func (p *Payment) Refund(id uint) error {
	tx := p.db.
		Model(&models.Payment{}).
		Where("id = ?", id).
		Update("status", models.PaymentStatusRefunded)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "dao.payment.refund")
	}

	return nil
}
