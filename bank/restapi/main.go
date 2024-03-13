package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"deuna.com/payment/bank/restapi/router"
)

var (
	port       = flag.Int("port", 8081, "Port to run the server on")
	authServer = flag.String("auth-server", "http://localhost:8080", "Auth server URL")
)

func main() {
	flag.Parse()

	logrus.SetLevel(logrus.DebugLevel)

	r := router.New(*authServer)

	host := fmt.Sprintf(":%d", *port)

	logrus.Info("banks running on ", host)
	logrus.Fatal(http.ListenAndServe(host, r))
}
