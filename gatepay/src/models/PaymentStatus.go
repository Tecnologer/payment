package models

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate enumer -type=PaymentStatus -json -transform=snake -sql -trimprefix=PaymentStatus
type PaymentStatus byte

const (
	PaymentStatusApproved PaymentStatus = iota
	PaymentStatusRefunded
)

func (ps PaymentStatus) GormDataType() string {
	return "text"
}

func (ps *PaymentStatus) GormValue(ctx context.Context, db *gorm.DB) (expr clause.Expr) {
	return clause.Expr{SQL: "?", Vars: []interface{}{ps.String()}}
}
