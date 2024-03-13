package router

import (
	"net/http"

	"deuna.com/payment/gatepay/restapi/handler"

	"github.com/gorilla/mux"
)

func New(h *handler.PaymentHandler) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/pay", h.Pay).Methods("POST")

	return router
}
