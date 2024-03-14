package router

import (
	"net/http"

	"deuna.com/payment/bank/constants"

	"deuna.com/payment/bank/restapi/handler"
	"github.com/gorilla/mux"
)

func New(authServer string) http.Handler {
	paymentHandler := handler.NewPaymentHandler(authServer)
	accountHandler := handler.NewAccountHandler(authServer)

	router := mux.NewRouter()

	router.HandleFunc(constants.TransferEndPoint, paymentHandler.Payment).Methods("POST")
	router.HandleFunc(constants.GetAccountEndPoint, accountHandler.GetAccount).Methods("POST")

	return router
}
