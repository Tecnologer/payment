package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"deuna.com/payment/gatepay/restapi/handler"
	"deuna.com/payment/gatepay/restapi/router"
)

var (
	port     = flag.Int("port", 8082, "Port")
	bankHost = flag.String("bank-server", "http://localhost:8081", "BankName host")
	authHost = flag.String("auth-server", "http://localhost:8080", "Auth host")
)

func main() {
	flag.Parse()

	logrus.SetLevel(logrus.DebugLevel)

	h := handler.NewPaymentHandler(*authHost, *bankHost)
	r := router.New(h)

	host := fmt.Sprintf(":%d", *port)

	logrus.Info("gatepay running on ", host)
	logrus.Fatal(http.ListenAndServe(host, r))
}
