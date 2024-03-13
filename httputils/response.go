package httputils

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func RetrieveError(r *http.Response) error {
	var res map[string]string

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		return errors.Wrap(err, "httputils.retrieve_error: decoding response")
	}

	return errors.New(res["error"])
}
