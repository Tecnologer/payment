package handler

import (
	"encoding/json"
	"net/http"

	"deuna.com/payment/bank/factory"

	"github.com/pkg/errors"

	"deuna.com/payment/httputils"

	"deuna.com/payment/bank/restapi/viewmodels"
)

type AccountHandler struct {
	*Handler
}

func NewAccountHandler(authHost string) *AccountHandler {
	return &AccountHandler{
		Handler: New(authHost),
	}
}

func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	isAuthorized, err := h.isTokenValid(r)
	if err != nil {
		httputils.WriteInternalServerError(w, errors.Wrap(err, "bank.account_handler.get_account: validating token"))

		return
	}

	if !isAuthorized {
		httputils.WriteUnauthorized(w, errors.New("bank.account_handler.get_account: invalid token"))

		return
	}

	var viewAccount *viewmodels.Account

	err = json.NewDecoder(r.Body).Decode(&viewAccount)
	if err != nil {
		httputils.WriteInternalServerError(w, errors.Wrap(err, "bank.account_handler.get_account: decoding request"))

		return
	}

	bank := factory.NewBank(viewAccount.BankName)
	if bank == nil {
		httputils.WriteBadRequest(w, errors.Errorf("bank not found: %s", viewAccount.BankName))

		return
	}

	account, err := bank.GetAccount(viewAccount.Number)
	if err != nil {
		httputils.WriteInternalServerError(w, errors.Wrap(err, "bank.account_handler.get_account: getting account"))

		return
	}

	httputils.WriteOK(w, account)
}
