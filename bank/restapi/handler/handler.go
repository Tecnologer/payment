package handler

import (
	"net/http"

	"deuna.com/payment/auth/api"
	"github.com/pkg/errors"
)

type Handler struct {
	Auth *api.Auth
}

func New(authHost string) *Handler {
	return &Handler{
		Auth: api.NewAuth(authHost),
	}
}

func (h *Handler) getTokenUser(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", errors.New("bank.handler.get_token_user: missing token")
	}

	claim, err := h.Auth.GetClaim(token)
	if err != nil {
		return "", errors.Wrap(err, "bank.handler.get_token_user: getting claim")
	}

	return claim.Username, nil
}

func (h *Handler) isTokenValid(r *http.Request) (bool, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return false, errors.New("bank.handler.is_token_valid: missing token")
	}

	isAuthorized, err := h.Auth.IsAuthorized(token)
	if err != nil {
		return false, errors.Wrap(err, "bank.handler.is_token_valid: validating token")
	}

	return isAuthorized, nil
}
