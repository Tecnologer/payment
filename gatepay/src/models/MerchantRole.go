package models

//go:generate enumer -type=MerchantRole -json -sql -transform=snake
type MerchantRole byte

const (
	Manager MerchantRole = iota
	Staff
	SuperAdmin
)
