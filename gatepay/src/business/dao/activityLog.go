package dao

import (
	"encoding/json"
	"fmt"

	"deuna.com/payment/gatepay/src/activityLog"

	"deuna.com/payment/gatepay/src/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ActivityLog struct {
	db *gorm.DB
}

func NewActivityLog(db *gorm.DB) *ActivityLog {
	return &ActivityLog{db}
}

func (a *ActivityLog) RegisterPayment(author string, payment *models.Payment) error {
	return a.registerNewPayMovement(author, models.ActivityLogTypePayment, payment)
}

func (a *ActivityLog) RegisterRefund(author string, payment *models.Payment) error {
	return a.registerNewPayMovement(author, models.ActivityLogTypeRefund, payment)
}

func (a *ActivityLog) RegisterNewPaymentMethod(paymentMethod *models.PaymentMethod) error {
	return a.registerPaymentMethod(paymentMethod.OwnerEmail, models.ActivityLogActionCreate, paymentMethod)
}

func (a *ActivityLog) Retrieve(pagination *activityLog.Pagination) (activityLogs []*models.ActivityLog, _ error) {
	db := a.db
	if pagination != nil {
		db = a.db.Scopes(pagination.Scopes()...)
	}

	tx := db.Find(&activityLogs)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "dao.activity_log.retrieve: failed to retrieve activity logs")
	}

	return
}

func (a *ActivityLog) registerNewPayMovement(author string, t models.ActivityLogType, payment *models.Payment) error {
	detail, err := a.buildPaymentDetail(payment)
	if err != nil {
		return errors.Wrap(err, "dao.activity_log.register_new_payment: failed to build payment detail")
	}

	log := a.logsFactory(t, author)
	log.Detail = detail

	tx := a.db.Save(log)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "dao.activity_log.register_new_payment: failed to save activity log")
	}

	return nil
}

func (a *ActivityLog) buildPaymentDetail(payment *models.Payment) (models.ActivityLogDetail, error) {
	paymentData := map[string]interface{}{
		"id":     payment.ID,
		"amount": payment.Amount,
	}

	return json.Marshal(paymentData)
}

func (a *ActivityLog) logsFactory(t models.ActivityLogType, author string) *models.ActivityLog {
	switch t {
	case models.ActivityLogTypePayment:
		return models.NewActivityLogPayment(author)
	case models.ActivityLogTypeRefund:
		return models.NewActivityLogRefund(author)
	case models.ActivityLogTypePaymentMethod:
		return models.NewActivityLogPaymentMethod(author, models.ActivityLogActionCreate)
	default:
		return nil
	}
}

func (a *ActivityLog) registerPaymentMethod(
	author string,
	action models.ActivityLogAction,
	paymentMethod *models.PaymentMethod,
) error {
	detail, err := a.buildPaymentMethodDetail(paymentMethod)
	if err != nil {
		return errors.Wrap(err, "dao.activity_log.register_payment_method: failed to build payment method detail")
	}

	log := models.NewActivityLogPaymentMethod(author, action)
	log.Detail = detail

	tx := a.db.Save(log)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "dao.activity_log.register_payment_method: failed to save activity log")
	}

	return nil
}

func (a *ActivityLog) buildPaymentMethodDetail(paymentMethod *models.PaymentMethod) (models.ActivityLogDetail, error) {
	paymentMethodData := map[string]interface{}{
		"id":      paymentMethod.ID,
		"name":    paymentMethod.Name,
		"account": fmt.Sprintf("%s-%s", paymentMethod.BankName, paymentMethod.AccountNumber),
	}

	return json.Marshal(paymentMethodData)
}
