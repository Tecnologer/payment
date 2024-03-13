package handler

import (
	"encoding/json"
	"net/http"

	"deuna.com/payment/httputils"

	"deuna.com/payment/bank/factory"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type PaymentHandler struct {
	*Handler
}

func NewPaymentHandler(authHost string) *PaymentHandler {
	return &PaymentHandler{
		Handler: New(authHost),
	}
}

type Payment struct {
	OriginBank         string  `json:"origin_bank"`
	OriginAccount      string  `json:"origin_account"`
	DestinationBank    string  `json:"destination_bank"`
	DestinationAccount string  `json:"destination_account"`
	Amount             float32 `json:"amount"`
}

type PaymentSummary struct {
	OriginBank         string  `json:"origin_bank"`
	OriginAccount      string  `json:"origin_account"`
	DestinationBank    string  `json:"destination_bank"`
	DestinationAccount string  `json:"destination_account"`
	Amount             float32 `json:"amount"`
	Status             string  `json:"status"`
	Error              string  `json:"error,omitempty"`
}

func (h *PaymentHandler) Payment(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("payment_handler_payment: start")

	username, err := h.getTokenUser(r)
	if err != nil {
		httputils.WriteUnauthorized(w, errors.Wrap(err, "getting token user"))

		return
	}

	var payment *Payment

	err = json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		httputils.WriteBadRequest(w, errors.Wrap(err, "decoding payment"))

		return
	}

	logrus.WithField("username", username).WithField("payment", payment).Debug("payment_handler_payment: payment")

	err = h.doPayment(payment)
	if err != nil {
		httputils.WriteInternalServerError(w, errors.Wrap(err, "doing payment"))

		return
	}

	httputils.WriteOK(w, createSuccessPaymentSummary(payment))
}

func (h *PaymentHandler) doPayment(payment *Payment) error {
	originBank := factory.NewBank(payment.OriginBank)
	if originBank == nil {
		return errors.Errorf("origin bank %s not found", payment.OriginBank)
	}

	destinationBank := factory.NewBank(payment.DestinationBank)
	if destinationBank == nil {
		return errors.Errorf("destination bank %s not found", payment.DestinationBank)
	}

	originAccount, err := originBank.GetAccount(payment.OriginAccount)
	if err != nil {
		return errors.Wrapf(err, "getting account %s", payment.OriginAccount)
	}

	destinationAccount, err := destinationBank.GetAccount(payment.DestinationAccount)
	if err != nil {
		return errors.Wrapf(err, "getting destination account %s", payment.DestinationAccount)
	}

	err = originBank.Payment(originAccount, destinationAccount, payment.Amount)
	if err != nil {
		return errors.Wrap(err, "doing payment")
	}

	return nil
}

func createSuccessPaymentSummary(payment *Payment) *PaymentSummary {
	return &PaymentSummary{
		OriginBank:         payment.OriginBank,
		OriginAccount:      payment.OriginAccount,
		DestinationBank:    payment.DestinationBank,
		DestinationAccount: payment.DestinationAccount,
		Amount:             payment.Amount,
		Status:             "success",
	}
}
