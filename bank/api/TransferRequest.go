package api

import (
	"context"
	"encoding/json"

	"deuna.com/payment/bank/models/interfaces"
)

type TransferRequest struct {
	Request
	Origin      interfaces.Account
	Destination interfaces.Account
	Amount      float32
}

func NewTransferRequest(ctx context.Context, origin, destination interfaces.Account, amount float32) *TransferRequest {
	return &TransferRequest{
		Request:     Request{Context: ctx},
		Origin:      origin,
		Destination: destination,
		Amount:      amount,
	}
}

func (t TransferRequest) MarshalJSON() ([]byte, error) {
	transferInfo := map[string]interface{}{
		"origin_bank":         t.Origin.GetBankName(),
		"origin_account":      t.Origin.GetID(),
		"destination_bank":    t.Destination.GetBankName(),
		"destination_account": t.Destination.GetID(),
		"amount":              t.Amount,
	}

	return json.Marshal(transferInfo)
}
