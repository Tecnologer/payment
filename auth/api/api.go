package api

import (
	"bytes"
	"deuna.com/payment/httputils"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"deuna.com/payment/auth/models"
	"github.com/pkg/errors"
)

type Auth struct {
	host string
}

func NewAuth(host string) *Auth {
	return &Auth{
		host: host,
	}
}

func (a *Auth) Login(username string, password string) (string, error) {
	credentials := &models.Credential{
		Username: username,
		Password: password,
	}

	credentialsJson, err := json.Marshal(credentials)
	if err != nil {
		return "", errors.Wrap(err, "marshalling credentials")
	}

	res, err := http.Post(a.host+"/login", "application/json", bytes.NewReader(credentialsJson))
	if err != nil {
		return "", errors.Wrap(err, "request login")
	}

	if res.StatusCode != http.StatusOK {
		return "", errors.Errorf("invalid status code: %d", res.StatusCode)
	}

	var token map[string]string

	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		return "", errors.Wrap(err, "decoding token")
	}

	return token["token"], nil
}

func (a *Auth) IsAuthorized(token string) (bool, error) {
	req, err := http.NewRequest("GET", a.host+"/validate", nil)
	if err != nil {
		return false, errors.Wrap(err, "creating request")
	}

	logrus.Debug("token: ", token)

	req.Header.Set("Authorization", token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, errors.Wrap(err, "request validate")
	}

	if res.StatusCode != http.StatusOK {
		return false, nil
	}

	return true, nil
}

func (a *Auth) GetClaim(token string) (*models.Claim, error) {
	req, err := http.NewRequest("GET", a.host+"/validate", nil)
	if err != nil {
		return nil, errors.Wrap(err, "auth_api.get_claim: creating request")
	}

	logrus.Debug("token: ", token)

	req.Header.Set("Authorization", token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "auth_api.get_claim: request validate")
	}

	if res.StatusCode != http.StatusOK {
		logrus.WithError(httputils.RetrieveError(res)).Error("auth_api.get_claim: invalid status code")

		return nil, nil
	}

	var claim *models.Claim

	err = json.NewDecoder(res.Body).Decode(&claim)
	if err != nil {
		return nil, errors.Wrap(err, "auth_api.get_claim: decoding claim")
	}

	return claim, nil
}
