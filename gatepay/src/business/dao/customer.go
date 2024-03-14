package dao

import (
	"deuna.com/payment/gatepay/src/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Customer struct {
	db *gorm.DB
}

func NewCustomer(db *gorm.DB) *Customer {
	return &Customer{db: db}
}

// FindByName returns a customer by its name
func (c *Customer) FindByName(name string) (*models.Customer, error) {
	var customer *models.Customer

	tx := c.db.
		Preload("PaymentHistory").
		Preload("PaymentMethods").
		Where("name = ?", name).
		First(&customer)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(tx.Error, "dao.customer.find_by_name: finding customer by name")
	}

	return customer, nil
}

func (c *Customer) FindByEmail(email string) (*models.Customer, error) {
	var customer *models.Customer

	tx := c.db.
		Preload("PaymentMethods").
		Where("email = ?", email).
		First(&customer)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(tx.Error, "dao.customer.find_by_email: finding customer by email")
	}

	return customer, nil
}

func (c *Customer) IsEmailCustomer(email string) (bool, error) {
	customer, err := c.FindByEmail(email)
	if err != nil {
		return false, errors.Wrap(err, "dao.customer.is_email_customer: finding customer by email")
	}

	return customer != nil, nil
}
