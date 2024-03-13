module deuna.com/payment/bank

go 1.22.0

replace (
	deuna.com/payment/auth => ../auth
	deuna.com/payment/httputils => ../httputils
)

require (
	deuna.com/payment/auth v0.0.0-00010101000000-000000000000
	deuna.com/payment/httputils v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)
