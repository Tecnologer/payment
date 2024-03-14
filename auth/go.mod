module deuna.com/payment/auth

go 1.22.0

require (
	deuna.com/payment/httputils v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.9.3
)

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect

replace deuna.com/payment/httputils => ../httputils
