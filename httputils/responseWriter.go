package httputils

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func WriteBadRequest(w http.ResponseWriter, err error) {
	WriteError(w, http.StatusBadRequest, err)
}

func WriteUnauthorized(w http.ResponseWriter, err error) {
	WriteError(w, http.StatusUnauthorized, err)
}

func WriteInternalServerError(w http.ResponseWriter, err error) {
	WriteError(w, http.StatusInternalServerError, err)
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	logrus.WithError(err).Debugf("writing error with status %s (%d)", http.StatusText(statusCode), statusCode)

	WriteHeaders(w, statusCode)

	_, err = w.Write(buildError(err))
	if err != nil {
		logrus.WithError(err).Error("writing error")
	}
}

func WriteOK(w http.ResponseWriter, response interface{}) {
	logrus.Debugf("success response")

	WriteHeaders(w, http.StatusOK)

	_, err := w.Write(buildResponse(response))
	if err != nil {
		logrus.WithError(err).Error("writing response")
	}
}

func WriteHeaders(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)
}

func buildError(err error) []byte {
	errResponse := map[string]string{
		"error": err.Error(),
	}

	response, _ := json.Marshal(errResponse)

	return response
}

func buildResponse(response interface{}) []byte {
	responseBytes, _ := json.Marshal(response)

	return responseBytes
}
