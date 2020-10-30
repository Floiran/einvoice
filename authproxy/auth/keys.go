package auth

import (
	"crypto/rsa"
	"strings"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	. "github.com/slovak-egov/einvoice/authproxy/config"
)

type Keys struct {
	ApiTokenPrivate *rsa.PrivateKey
	OboTokenPublic  *rsa.PublicKey
}

func NewKeys() *Keys {
	signBytes := "-----BEGIN RSA PRIVATE KEY-----\n" +
		strings.ReplaceAll(Config.SlovenskoSk.ApiTokenPrivateKey, " ", string(byte(10))) +
		"\n-----END RSA PRIVATE KEY-----\n"
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(signBytes))
	if err != nil {
		log.WithField("error", err).Fatal("new_keys.parse_rsa_private_error")
	}

	verifyBytes := "-----BEGIN RSA PRIVATE KEY-----\n" +
		strings.ReplaceAll(Config.SlovenskoSk.OboTokenPublicKey, " ", string(byte(10))) +
		"\n-----END RSA PRIVATE KEY-----\n"
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(verifyBytes))
	if err != nil {
		log.WithField("error", err).Fatal("new_keys.parse_rsa_public_error")
	}

	return &Keys{
		ApiTokenPrivate: signKey,
		OboTokenPublic:  verifyKey,
	}
}
