package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"github.com/slovak-egov/einvoice/authproxy/db"
)

func (a *App) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")

		userId, err := a.manager.GetUserIdByToken(token)
		if err != nil {
			userId, err = a.getUserByServiceAccount(token)
		}

		if err != nil {
			res.WriteHeader(401)
			return
		}

		req.Header.Del("Authorization")
		req.Header.Set("User-Id", userId)

		// Call the next handler
		next.ServeHTTP(res, req)
	})
}


func (a *App) getUserByServiceAccount(tokenString string) (string, error) {
	var user *db.User
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("Cannot parse claims")
		}

		user, _ = a.manager.GetUser(claims["sub"].(string))
		if user == nil {
			return nil, errors.New("User not found")
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(user.ServiceAccountKey))
		if err != nil {
			return nil, errors.New("Invalid key")
		}

		return verifyKey, nil
	})

	if err != nil {
		log.WithField("token", tokenString).Debug("app.auth_middleware.parse_token.failed")
		return "", err
	}

	if !token.Valid {
		log.WithField("token", tokenString).Debug("authenticator.token.verify.failed")
		return "", errors.New("Invalid token")
	}

	return user.Id, nil
}
