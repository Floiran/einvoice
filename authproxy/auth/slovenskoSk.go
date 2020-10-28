package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/slovak-egov/einvoice/authproxy/config"
	"github.com/slovak-egov/einvoice/common"
	"io/ioutil"
	"net/http"
)

type SlovenskoSkUser struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
}

func GetSlovenkosSkUserInfo(keys *Keys, oboToken string) (*SlovenskoSkUser, error) {
	token, err := jwt.Parse(oboToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return keys.OboTokenPublic, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Cannot parse claims")
	}

	slovenskoSkToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp": claims["exp"],
		"jti": common.RandomString(32),
		"obo": oboToken,
	})
	slovenskoSkToken.Header["alg"] = "RS256"
	slovenskoSkToken.Header["cty"] = "JWT"
	delete(slovenskoSkToken.Header, "typ")

	slovenskoSkTokenString, err := slovenskoSkToken.SignedString(keys.ApiTokenPrivate)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	slovenskoSkReq, err := http.NewRequest("GET", config.Config.SlovenskoSk.Url+"/api/upvs/user/info", nil)
	if err != nil {
		return nil, err
	}
	slovenskoSkReq.Header.Add("Authorization", "Bearer "+slovenskoSkTokenString)
	slovenskoSkRes, err := client.Do(slovenskoSkReq)
	if err != nil {
		return nil, err
	}

	defer slovenskoSkRes.Body.Close()
	body, err := ioutil.ReadAll(slovenskoSkRes.Body)
	if err != nil {
		return nil, err
	}

	user := &SlovenskoSkUser{}
	if err = json.Unmarshal(body, user); err != nil {
		return nil, err
	}

	if user.Uri == "" {
		return nil, errors.New("Unauthorized")
	}

	return user, nil
}
