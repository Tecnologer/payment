package db

//go:generate enumer -type=SSLMode -json -sql -transform=snake
type SSLMode byte

const (
	Prefer SSLMode = iota
	Disable
	Require
	Allow
	Verify
	VerifyFull
)
