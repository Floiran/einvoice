package auth

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/slovak-egov/einvoice/authproxy/user"
)

type Authenticator interface {
	WithUser(func(res http.ResponseWriter, req *http.Request, userId string)) func(http.ResponseWriter, *http.Request)
}

type authenticator struct {
	userManager UserManager
}

func NewAuthenticator(userManager UserManager) Authenticator {
	return &authenticator{userManager: userManager}
}

func (auth *authenticator) WithUser(f func(res http.ResponseWriter, req *http.Request, userId string)) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")

		userId, err := auth.userManager.GetUserIdByToken(token)
		if err != nil {
			userId, err = auth.userByServiceAccount(token)
		}

		if err == nil {
			req.Header.Del("Authorization")
			f(res, req, userId)
		} else {
			res.WriteHeader(401)
		}
	}
}

func (auth *authenticator) userByServiceAccount(tokenString string) (string, error) {
	var user *user.User
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("Cannot parse claims")
		}

		user, _ = auth.userManager.GetUser(claims["sub"].(string))
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
		log.WithFields(log.Fields{
			"error": err.Error(),
			"token": tokenString,
		}).Debug("authenticator.user_by_service_account.token_parsing.failed")

		return "", err
	}

	if !token.Valid {
		log.WithField("token", tokenString).Debug("authenticator.token.verify.failed")
		return "", errors.New("Invalid token")
	}

	return user.Id, nil
}
