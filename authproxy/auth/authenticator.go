package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/slovak-egov/einvoice/authproxy/user"
)

type Authenticator interface {
	WithUser(func(http.ResponseWriter, *http.Request, *user.User)) func(http.ResponseWriter, *http.Request)
}

type authenticator struct {
	userManager UserManager
}

func NewAuthenticator(userManager UserManager) Authenticator {
	return &authenticator{userManager: userManager}
}

func (auth *authenticator) WithUser(f func(http.ResponseWriter, *http.Request, *user.User)) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")

		user := auth.userManager.GetUserByToken(token)
		if user == nil {
			user = auth.verifyJwtToken(token)
		}

		if user != nil {
			req.Header.Del("Authorization")
			f(res, req, user)
		} else {
			res.WriteHeader(401)
		}
	}
}

func (auth *authenticator) verifyJwtToken(tokenString string) *user.User {
	var user *user.User
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("Cannot parse claims")
		}

		user = auth.userManager.GetUser(claims["sub"].(string))
		if user == nil {
			return nil, errors.New("User not found")
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(user.ServiceAccountKey))
		if err != nil {
			return nil, errors.New("Invalid key")
		}

		return verifyKey, nil
	})

	if err != nil || !token.Valid {
		return nil
	}

	return user
}
