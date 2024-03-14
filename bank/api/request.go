package api

import (
	"context"
	"net/http"

	"deuna.com/payment/httputils"
)

type Request struct {
	Context context.Context `json:"-"`
}

func (r Request) PrepareRequestHeaders(req *http.Request) {
	req.Header.Set("content-type", "application/json")
	r.WriteTokenToRequest(req)
}

func (r Request) WriteTokenToRequest(w *http.Request) {
	w.Header.Set("Authorization", r.Context.Value(httputils.TokenKey).(string))
}
