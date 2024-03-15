package models

//go:generate enumer -type=ActivityLogType -json -transform=snake -trimprefix=ActivityLogType -sql
type ActivityLogType byte

const (
	ActivityLogTypeNone ActivityLogType = iota
	ActivityLogTypePayment
	ActivityLogTypeRefund
	ActivityLogTypePaymentMethod
)

//go:generate enumer -type=ActivityLogAction -json -transform=snake -trimprefix=ActivityLogAction -sql
type ActivityLogAction byte

const (
	ActivityLogActionNone ActivityLogAction = iota
	ActivityLogActionCreate
	ActivityLogActionUpdate
	ActivityLogActionDelete
)
