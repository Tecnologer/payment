package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"deuna.com/payment/httputils"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"

	"deuna.com/payment/auth/models"
)

var (
	port = flag.Int("port", 8080, "Port to run the server on")

	jwtSecret = "$3cr3t#"
)

func main() {
	flag.Parse()

	logrus.SetLevel(logrus.DebugLevel)

	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		jwtSecret = secret
	}

	http.HandleFunc("/login", login)
	http.HandleFunc("/validate", validate)

	host := fmt.Sprintf(":%d", *port)

	logrus.Info("Server running on ", host)
	logrus.Fatal(http.ListenAndServe(host, nil))
}

// signToken handler to authenticate users and generate JWT
func login(w http.ResponseWriter, r *http.Request) {
	var creds models.Credential

	logrus.Debug("signToken request for user: ", creds.Username)

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString, err := signToken(&creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.WithError(err).Error("getting token")

		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		logrus.WithError(err).Error("Error encoding token")

		return
	}

	w.WriteHeader(http.StatusOK)
	logrus.Debugf("User %s logged in successfully", creds.Username)
}

func signToken(credentials *models.Credential) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claim{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", errors.Wrap(err, "auth_api_login: signing token")
	}

	return tokenString, nil
}

func validate(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	logrus.Debug("validate request for token: ", token)

	claims, err := validateToken(token)
	if err != nil {
		logrus.WithError(err).Error("auth_api_validate: validating token")

		httputils.WriteUnauthorized(w, errors.Wrap(err, "auth_api_validate: validating token"))

		return
	}

	w.Header().Set("Content-Type", "application/json")

	httputils.WriteOK(w, claims)
}

func validateToken(tokenString string) (*models.Claim, error) {
	claims := &models.Claim{}

	// Convert your jwtSecret to a byte slice if it's stored as a string
	secret := []byte(jwtSecret)

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "auth_api_validate_token: parsing token")
	}

	if !token.Valid {
		return nil, errors.New("auth_api_validate_token: invalid token")
	}

	return claims, nil
}
