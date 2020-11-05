package app

import (
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(req *http.Request) (string, error) {
	tokenHeader := req.Header.Get("Authorization")
	if tokenHeader == "" {
		return "", errors.New("Missing authorization")
	}

	parts := strings.Split(tokenHeader, " ")
	if len(parts) != 2 {
		return "", errors.New("Invalid token format")
	}
	if parts[0] != "Bearer" {
		return "", errors.New("Invalid authorization type")
	}
	return parts[1], nil
}
