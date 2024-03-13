package router

import (
	"net/http"

	"deuna.com/payment/bank/restapi/handler"
	"github.com/gorilla/mux"
)

func New(authServer string) http.Handler {
	paymentHandler := handler.NewPaymentHandler(authServer)
	accountHandler := handler.NewAccountHandler(authServer)

	router := mux.NewRouter()

	router.HandleFunc("/payment", paymentHandler.Payment).Methods("POST")
	router.HandleFunc("/get-account", accountHandler.GetAccount).Methods("POST")

	return router
}
