package slovenskoSk

import (
	"crypto/rsa"
	"strings"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"github.com/slovak-egov/einvoice/authproxy/config"
)

type Connector struct {
	baseUrl string
	apiTokenPrivate *rsa.PrivateKey
	oboTokenPublic  *rsa.PublicKey
}

func Init(config config.SlovenskoSkConfiguration) Connector {
	return Connector{
		baseUrl: config.Url,
		apiTokenPrivate: getPrivateKey(config.ApiTokenPrivateKey),
		oboTokenPublic: getPublicKey(config.OboTokenPublicKey),
	}
}

func getPrivateKey(privateKey string) *rsa.PrivateKey {
	signBytes := "-----BEGIN RSA PRIVATE KEY-----\n" +
		strings.ReplaceAll(privateKey, " ", string(byte(10))) +
		"\n-----END RSA PRIVATE KEY-----\n"
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(signBytes))
	if err != nil {
		log.WithField("error", err.Error()).Fatal("slovenskosk.keys.parse_private")
	}

	return signKey
}

func getPublicKey(publicKey string) *rsa.PublicKey {
	verifyBytes := "-----BEGIN RSA PRIVATE KEY-----\n" +
		strings.ReplaceAll(publicKey, " ", string(byte(10))) +
		"\n-----END RSA PRIVATE KEY-----\n"
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(verifyBytes))
	if err != nil {
		log.WithField("error", err).Fatal("slovenskosk.keys.parse_public")
	}

	return verifyKey
}
