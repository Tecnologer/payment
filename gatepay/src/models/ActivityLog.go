package models

import (
	"deuna.com/payment/gatepay/src/validator"
	"gorm.io/gorm"
)

type ActivityLog struct {
	*gorm.Model
	Type   ActivityLogType   `json:"type"   validate:"gte=0"`
	Author string            `json:"author" validate:"required,email"`
	Action ActivityLogAction `json:"action" validate:"gte=0"`
	Detail ActivityLogDetail `json:"detail"`
}

func (l *ActivityLog) Validate() error {
	return validator.ValidateObject(l)
}

func NewActivityLogPayment(author string) *ActivityLog {
	return &ActivityLog{
		Type:   ActivityLogTypePayment,
		Author: author,
		Action: ActivityLogActionCreate,
	}
}

func NewActivityLogRefund(author string) *ActivityLog {
	return &ActivityLog{
		Type:   ActivityLogTypeRefund,
		Author: author,
		Action: ActivityLogActionCreate,
	}
}

func NewActivityLogPaymentMethod(author string, action ActivityLogAction) *ActivityLog {
	return &ActivityLog{
		Type:   ActivityLogTypePaymentMethod,
		Author: author,
		Action: action,
	}
}
