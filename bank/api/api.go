package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"deuna.com/payment/bank/constants"

	"deuna.com/payment/httputils"

	"deuna.com/payment/bank/models"

	"deuna.com/payment/bank/models/interfaces"
	"github.com/pkg/errors"
)

type BankService struct {
	router
}

func NewBankService(host string) *BankService {
	return &BankService{
		router: newRouter(host),
	}
}

func (b *BankService) GetAccount(request *AccountRequest) (interfaces.Account, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, errors.Wrap(err, "api.bank_service.get_account: marshaling account info")
	}

	req, err := http.NewRequest(http.MethodPost, b.buildPath(constants.GetAccountEndPoint), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, errors.Wrap(err, "api.bank_service.get_account: creating request")
	}

	request.PrepareRequestHeaders(req)

	httpClient := &http.Client{}

	res, err := httpClient.Do(req)
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

func (b *BankService) Transfer(request *TransferRequest) error {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return errors.Wrap(err, "api.bank_service.transfer: marshaling transfer info")
	}

	req, err := http.NewRequest(http.MethodPost, b.buildPath(constants.TransferEndPoint), bytes.NewBuffer(reqBody))
	if err != nil {
		return errors.Wrap(err, "api.bank_service.get_account: creating request")
	}

	request.PrepareRequestHeaders(req)

	httpClient := &http.Client{}

	res, err := httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "api.bank_service.get_account: calling bank service")
	}

	if res.StatusCode != http.StatusOK {
		return httputils.RetrieveError(res)
	}

	return nil
}
