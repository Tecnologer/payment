package dao

import (
	"deuna.com/payment/gatepay/src/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	db *gorm.DB
}

func NewPaymentMethod(db *gorm.DB) *PaymentMethod {
	return &PaymentMethod{db: db}
}

func (p *PaymentMethod) Insert(inputPaymentMethod *models.PaymentMethod) (*models.PaymentMethod, error) {
	tx := p.db.Save(inputPaymentMethod)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return inputPaymentMethod, nil
}

func (p *PaymentMethod) Update(paymentMethod *models.PaymentMethod) (*models.PaymentMethod, error) {
	tx := p.db.Save(paymentMethod)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return paymentMethod, nil
}

func (p *PaymentMethod) InsertIfNoExists(inputPaymentMethod *models.PaymentMethod) (*models.PaymentMethod, error) {
	paymentMethod, err := p.FindByBankNAccount(inputPaymentMethod.BankName, inputPaymentMethod.AccountNumber)
	if err != nil {
		return nil, errors.Wrap(err, "dao.payment_method.insert_if_no_exists: finding payment method")
	}

	if paymentMethod == nil {
		paymentMethod, err = p.Insert(inputPaymentMethod)
		if err != nil {
			return nil, errors.Wrap(err, "dao.payment_method.insert_if_no_exists: inserting payment method")
		}
	}

	paymentMethod, err = p.FindByID(paymentMethod.ID)
	if err != nil {
		return nil, errors.Wrap(err, "dao.payment_method.insert_if_no_exists: retrieving saved payment method")
	}

	return paymentMethod, nil
}

func (p *PaymentMethod) FindByBankNAccount(bankName string, accountNumber string) (*models.PaymentMethod, error) {
	var paymentMethod *models.PaymentMethod

	tx := p.db.
		Preload("Customer").
		Preload("Merchant").
		Where("bank_name = ?", bankName).
		Where("account_number = ?", accountNumber).
		First(&paymentMethod)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(tx.Error, "dao.payment_method.find_by_bank_n_account")
	}

	return paymentMethod, nil
}

func (p *PaymentMethod) FindByID(id uint) (*models.PaymentMethod, error) {
	var paymentMethod *models.PaymentMethod

	tx := p.db.
		Preload("Customer").
		Preload("Merchant").
		First(&paymentMethod, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(tx.Error, "dao.payment_method.find_by_id")
	}

	return paymentMethod, nil
}
