package handler

import (
	"encoding/json"
	"net/http"

	"deuna.com/payment/gatepay/src/business"
	"deuna.com/payment/gatepay/src/db"

	"deuna.com/payment/auth/api"
	"deuna.com/payment/gatepay/restapi/viewmodels"
	"github.com/sirupsen/logrus"
)

type PaymentHandler struct {
	Auth *api.Auth
}

func NewPaymentHandler(authHost string) *PaymentHandler {
	return &PaymentHandler{
		Auth: api.NewAuth(authHost),
	}
}

func (h *PaymentHandler) Pay(w http.ResponseWriter, r *http.Request) {
	var viewPayment *viewmodels.Payment

	err := json.NewDecoder(r.Body).Decode(&viewPayment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	cnn, err := db.DefaultConnection()
	if err != nil {
		logrus.WithError(err).Error("handler.payment.pay: getting connection")
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	// Begin transaction
	cnn = cnn.Begin()

	inputPayment := viewPayment.ToPaymentModel()

	payment, err := business.NewPayment(cnn).Register(inputPayment)
	if err != nil {
		logrus.WithError(err).Error("handler.payment.pay: registering payment")
		w.WriteHeader(http.StatusInternalServerError)

		// Rollback transaction
		cnn.Rollback()

		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(payment)
	if err != nil {
		logrus.WithError(err).Error("handler.payment.pay: encoding response")
		w.WriteHeader(http.StatusInternalServerError)

		// Rollback transaction
		cnn.Rollback()

		return
	}

	w.WriteHeader(http.StatusOK)

	// Commit transaction
	cnn.Commit()

	logrus.Debugf(
		"handler.payment.pay: payment success: from %s to %s a total of %f",
		payment.OriginPaymentMethod.Name,
		payment.DestinationPaymentMethod.Name,
		payment.Amount,
	)
}
