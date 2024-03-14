package api

import (
	"net/http"

	"github.com/pkg/errors"
)

type AuthHandler struct {
	Auth *Auth
}

func NewAuthHandler(authHost string) *AuthHandler {
	return &AuthHandler{
		Auth: NewAuth(authHost),
	}
}

func (h *AuthHandler) EmailFromToken(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", errors.New("bank.handler.get_token_user: missing token")
	}

	claim, err := h.Auth.GetClaim(token)
	if err != nil {
		return "", errors.Wrap(err, "bank.handler.get_token_user: getting claim")
	}

	if claim == nil {
		return "", errors.New("bank.handler.get_token_user: unauthorized token")
	}

	return claim.Username, nil
}

func (h *AuthHandler) IsTokenValid(r *http.Request) (bool, error) {
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
