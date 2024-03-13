package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"deuna.com/payment/httputils"

	"deuna.com/payment/bank/models"

	"deuna.com/payment/bank/models/interfaces"
	"github.com/pkg/errors"
)

type BankService struct {
	host string
}

func NewBankService(host string) *BankService {
	return &BankService{
		host: host,
	}
}

func (b *BankService) GetAccount(bankName, accountNumber string) (interfaces.Account, error) {
	accountInfo := map[string]string{
		"bank_name": bankName,
		"number":    accountNumber,
	}

	req, err := json.Marshal(accountInfo)
	if err != nil {
		return nil, errors.Wrap(err, "api.bank_service.get_account: marshaling account info")
	}

	// Call to bank service
	res, err := http.Post(b.host+"/get-account", "application/json", bytes.NewBuffer(req))
	if err != nil {
		return nil, errors.Wrap(err, "api.bank_service.get_account: calling bank service")
	}

	if res.StatusCode != http.StatusOK {
		return nil, httputils.RetrieveError(res)
	}

	var account *models.Account

	err = json.NewDecoder(res.Body).Decode(&account)
	if err != nil {
		return nil, errors.Wrap(err, "api.bank_service.get_account: decoding response")
	}

	return account, nil
}
