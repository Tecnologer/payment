package business

import (
	"context"

	"deuna.com/payment/gatepay/src/activityLog"
	"deuna.com/payment/gatepay/src/business/dao"
	"deuna.com/payment/gatepay/src/models"
	"gorm.io/gorm"
)

type ActivityLog struct {
	*Business
	dao *dao.ActivityLog
}

func NewActivityLog(db *gorm.DB, ctx context.Context) *ActivityLog {
	return &ActivityLog{
		Business: NewBusiness(ctx),
		dao:      dao.NewActivityLog(db),
	}
}

func (a *ActivityLog) Retrieve(pagination *activityLog.Pagination) (activityLogs []*models.ActivityLog, _ error) {
	return a.dao.Retrieve(pagination)
}
