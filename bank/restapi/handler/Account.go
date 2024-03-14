package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"deuna.com/payment/auth/api"

	"deuna.com/payment/bank/factory"

	"github.com/pkg/errors"

	"deuna.com/payment/httputils"

	"deuna.com/payment/bank/restapi/viewmodels"
)

type AccountHandler struct {
	*api.AuthHandler
}

func NewAccountHandler(authHost string) *AccountHandler {
	return &AccountHandler{
		AuthHandler: api.NewAuthHandler(authHost),
	}
}

func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	isAuthorized, err := h.IsTokenValid(r)
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

	if err = viewAccount.Validate(); err != nil {
		httputils.WriteBadRequest(w, errors.Wrap(err, "bank.account_handler.get_account: validating account"))

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

	if account.GetOwnerName() != viewAccount.OwnerName {
		logrus.Infof("the account exists but the owner name is different: %s", viewAccount.OwnerName)

		httputils.WriteBadRequest(
			w,
			errors.Errorf(
				"there is not account with number %s and owner %s for bank %s",
				viewAccount.Number,
				viewAccount.OwnerName,
				viewAccount.BankName,
			),
		)

		return
	}

	httputils.WriteOK(w, account)
}
