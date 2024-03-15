package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"deuna.com/payment/gatepay/src/activityLog"

	"github.com/gorilla/mux"

	"deuna.com/payment/gatepay/restapi/viewmodels"

	"deuna.com/payment/httputils"
	"github.com/pkg/errors"

	"deuna.com/payment/gatepay/src/models"

	"deuna.com/payment/gatepay/src/business"
	"deuna.com/payment/gatepay/src/db"

	authApi "deuna.com/payment/auth/api"
	"github.com/sirupsen/logrus"
)

type PaymentHandler struct {
	*authApi.AuthHandler
}

func NewPaymentHandler(authHost, bankHost string) *PaymentHandler {
	// set environment variable for bank server url, used in service.Banker
	err := os.Setenv("BANK_SERVER_URL", bankHost)
	if err != nil {
		logrus.WithError(err).Error("handler.payment.new_payment_handler: setting bank server url")
	}

	return &PaymentHandler{
		AuthHandler: authApi.NewAuthHandler(authHost),
	}
}

func (h *PaymentHandler) Pay(w http.ResponseWriter, r *http.Request) {
	var inputPayment *viewmodels.Payment

	err := json.NewDecoder(r.Body).Decode(&inputPayment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	cnn, err := db.DefaultConnection()
	if err != nil {
		httputils.WriteInternalServerError(w, errors.Wrap(err, "getting connection"))

		return
	}

	ctx := httputils.ContextWithToken(r.Context(), r.Header.Get("Authorization"))

	senderEmail, err := h.EmailFromToken(r)
	if err != nil {
		logrus.WithError(err).Error("handler.payment.pay: getting token user")
		httputils.WriteUnauthorized(w, errors.Wrap(err, "getting token user"))

		return
	}

	newPayment := inputPayment.ParseToModelPayment()
	if err := newPayment.Validate(); err != nil {
		httputils.WriteBadRequest(w, errors.Wrap(err, "validating payment"))

		return
	}

	// Begin transaction
	cnn = cnn.Begin()

	payment, err := business.NewPayment(cnn, ctx).Register(senderEmail, newPayment)
	if err != nil {
		logrus.WithError(err).Error("handler.payment.pay: registering payment")
		httputils.WriteInternalServerError(w, errors.Wrap(err, "registering payment"))

		// Rollback transaction
		cnn.Rollback()

		return
	}

	httputils.WriteOK(w, payment)

	// Commit transaction
	cnn.Commit()

	logrus.Debugf(
		"handler.payment.pay: payment success: from %s to %s a total of %f",
		payment.OriginPaymentMethod.Name,
		payment.DestinationPaymentMethod.Name,
		payment.Amount,
	)
}

func (h *PaymentHandler) AddPaymentMethod(w http.ResponseWriter, r *http.Request) {
	var viewPaymentMethod *models.PaymentMethod

	err := json.NewDecoder(r.Body).Decode(&viewPaymentMethod)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	cnn, err := db.DefaultConnection()
	if err != nil {
		logrus.WithError(err).Error("handler.payment.add_payment_method: getting connection")
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	ctx := httputils.ContextWithToken(r.Context(), r.Header.Get("Authorization"))

	viewPaymentMethod.OwnerEmail, err = h.EmailFromToken(r)
	if err != nil {
		logrus.WithError(err).Error("handler.payment.add_payment_method: getting token user")

		httputils.WriteInternalServerError(w, errors.Wrap(err, "getting token user"))

		// Rollback transaction
		cnn.Rollback()

		return
	}

	// Begin transaction
	cnn = cnn.Begin()

	paymentMethod, err := business.NewPaymentMethod(cnn, ctx).Create(viewPaymentMethod)
	if err != nil {
		logrus.WithError(err).Error("handler.payment.add_payment_method: inserting payment method")
		httputils.WriteInternalServerError(w, errors.Wrap(err, "inserting payment method"))

		// Rollback transaction
		cnn.Rollback()

		return
	}

	httputils.WriteOK(w, paymentMethod)

	// Commit transaction
	cnn.Commit()

	logrus.Debugf(
		"handler.payment.add_payment_method: payment method %s added",
		paymentMethod.Name,
	)
}

func (h *PaymentHandler) GetPayments(w http.ResponseWriter, r *http.Request) {
	cnn, err := db.DefaultConnection()
	if err != nil {
		logrus.WithError(err).Error("handler.payment.get_payment_methods: getting connection")
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	ctx := httputils.ContextWithToken(r.Context(), r.Header.Get("Authorization"))

	ownerEmail, err := h.EmailFromToken(r)
	if err != nil {
		logrus.WithError(err).Error("handler.payment.get_payment_methods: getting token user")
		httputils.WriteInternalServerError(w, errors.Wrap(err, "getting token user"))

		return
	}

	paymentMethods, err := business.NewPayment(cnn, ctx).PaymentsByEmailOwner(ownerEmail)
	if err != nil {
		logrus.WithError(err).Error("handler.payment.get_payment_methods: getting payment methods")
		httputils.WriteInternalServerError(w, errors.Wrap(err, "getting payment methods"))

		return
	}

	httputils.WriteOK(w, paymentMethods)
}

func (h *PaymentHandler) RefundPayment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		logrus.WithError(err).Error("handler.payment.refund_payment: getting payment id")
		httputils.WriteBadRequest(w, errors.Wrap(err, "the id is not a number"))

		return

	}

	cnn, err := db.DefaultConnection()
	if err != nil {
		logrus.WithError(err).Error("handler.payment.refund_payment: getting connection")
		httputils.WriteInternalServerError(w, errors.Wrap(err, "getting connection"))

		return
	}

	ctx := httputils.ContextWithToken(r.Context(), r.Header.Get("Authorization"))

	userEmail, err := h.EmailFromToken(r)
	if err != nil {
		logrus.WithError(err).Error("handler.payment.refund_payment: getting token user")
		httputils.WriteUnauthorized(w, errors.Wrap(err, "getting token user"))

		return
	}

	// Begin transaction
	cnn = cnn.Begin()

	err = business.NewPayment(cnn, ctx).Refund(userEmail, uint(id))
	if err != nil {
		logrus.WithError(err).Error("handler.payment.refund_payment: refunding payment")
		httputils.WriteInternalServerError(w, errors.Wrap(err, "refunding payment"))

		// Rollback transaction
		cnn.Rollback()

		return
	}

	// Commit transaction
	cnn.Commit()

	httputils.WriteOK(w, nil)
}

func (h *PaymentHandler) GetActivityLog(w http.ResponseWriter, r *http.Request) {
	var paginationRequest *activityLog.Pagination

	err := json.NewDecoder(r.Body).Decode(&paginationRequest)
	if err != nil {
		logrus.WithError(err).Error("handler.payment.get_activity_log: decoding pagination request")
		httputils.WriteBadRequest(w, errors.Wrap(err, "decoding pagination request"))

		return
	}

	cnn, err := db.DefaultConnection()
	if err != nil {
		logrus.WithError(err).Error("handler.payment.get_activity_log: getting connection")
		httputils.WriteInternalServerError(w, errors.Wrap(err, "getting connection"))

		return
	}

	ctx := httputils.ContextWithToken(r.Context(), r.Header.Get("Authorization"))

	userEmail, err := h.EmailFromToken(r)
	if err != nil {
		logrus.WithError(err).Error("handler.payment.get_activity_log: getting token user")
		httputils.WriteUnauthorized(w, errors.Wrap(err, "getting token user"))

		return
	}

	paginationRequest.Filters = append(paginationRequest.Filters, activityLog.FilterByUserEmail(userEmail))

	activityLogs, err := business.NewActivityLog(cnn, ctx).Retrieve(paginationRequest)
	if err != nil {
		logrus.WithError(err).Error("handler.payment.get_activity_log: getting activity log")
		httputils.WriteInternalServerError(w, errors.Wrap(err, "getting activity log"))

		return
	}

	httputils.WriteOK(w, activityLogs)
}
