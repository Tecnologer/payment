package router

import (
	"net/http"

	"deuna.com/payment/gatepay/constants"

	"deuna.com/payment/gatepay/restapi/handler"

	"github.com/gorilla/mux"
)

func New(h *handler.PaymentHandler) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc(constants.PayEndPoint, h.Pay).Methods("POST")
	router.HandleFunc(constants.AddPaymentMethodEndPoint, h.AddPaymentMethod).Methods("POST")
	router.HandleFunc(constants.GetPaymentsEndPoint, h.GetPayments).Methods("GET")
	router.HandleFunc(constants.RefundPaymentEndPoint, h.RefundPayment).Methods("PUT")
	router.HandleFunc(constants.GetActivityLogEndPoint, h.GetActivityLog).Methods("POST")

	return router
}
