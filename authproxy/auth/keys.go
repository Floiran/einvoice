package auth

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/slovak-egov/einvoice/authproxy/config"
	"strings"
)

type Keys struct {
	ApiTokenPrivate *rsa.PrivateKey
	OboTokenPublic  *rsa.PublicKey
}

func NewKeys() *Keys {
	signBytes := "-----BEGIN RSA PRIVATE KEY-----\n" +
		strings.ReplaceAll(config.Config.SlovenskoSk.ApiTokenPrivateKey, " ", string(byte(10))) +
		"\n-----END RSA PRIVATE KEY-----\n"
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(signBytes))
	if err != nil {
		panic(err)
	}

	verifyBytes := "-----BEGIN RSA PRIVATE KEY-----\n" +
		strings.ReplaceAll(config.Config.SlovenskoSk.OboTokenPublicKey, " ", string(byte(10))) +
		"\n-----END RSA PRIVATE KEY-----\n"
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(verifyBytes))
	if err != nil {
		panic(err)
	}

	return &Keys{
		ApiTokenPrivate: signKey,
		OboTokenPublic:  verifyKey,
	}
}
