package auth

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	log "github.com/sirupsen/logrus"

	"github.com/slovak-egov/einvoice/authproxy/user"
)

type UserInfo struct {
	Token             string `json:"token,omitempty"`
	Id                string `json:"id"`
	Name              string `json:"name"`
	ServiceAccountKey string `json:"serviceAccountKey,omitempty"`
	Email             string `json:"email,omitempty"`
}

func userInfo(user *user.User) *UserInfo {
	return &UserInfo{
		Token:             user.Token,
		Id:                user.Id,
		Name:              user.Name,
		ServiceAccountKey: user.ServiceAccountKey,
		Email:             user.Email,
	}
}

func HandleLogin(manager UserManager, keys *Keys) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		tokenString := req.Header.Get("Authorization")
		slovenskoSkUser, err := GetSlovenskoSkUserInfo(keys, tokenString)
		if err != nil {
			log.WithField("error", err).Error("login.slovensko_sk.get_user_info")
			res.WriteHeader(401)
			return
		}

		id := slovenskoSkUser.Uri

		user := manager.GetUser(id)
		if user == nil {
			name := slovenskoSkUser.Name
			user = manager.Create(id, name)
		}

		manager.CreateToken(user)

		info := userInfo(user)
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(info)
	}
}

func HandleLogout(manager UserManager) func(res http.ResponseWriter, req *http.Request, user *user.User) {
	return func(res http.ResponseWriter, req *http.Request, user *user.User) {
		if err := manager.RemoveToken(user); err != nil {
			res.WriteHeader(401)
			return
		}

		res.WriteHeader(200)
	}
}

func HandleMe(res http.ResponseWriter, req *http.Request, user *user.User) {
	info := userInfo(user)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(info)
}

func HandleUpdateUser(manager UserManager) func(res http.ResponseWriter, req *http.Request, user *user.User) {
	return func(res http.ResponseWriter, req *http.Request, usr *user.User) {
		body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := req.Body.Close(); err != nil {
			panic(err)
		}

		updates := &user.User{}

		if err := json.Unmarshal(body, &updates); err != nil {
			panic(err)
		}

		manager.UpdateUser(usr, updates)

		info := userInfo(usr)
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(info)
	}
}

func HandleAuthProxy(proxy *httputil.ReverseProxy) func(res http.ResponseWriter, req *http.Request, user *user.User) {
	return func(res http.ResponseWriter, req *http.Request, user *user.User) {
		req.Host = req.URL.Host
		proxy.ServeHTTP(res, req)
	}
}

func HandleOpenProxy(proxy *httputil.ReverseProxy) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host
		proxy.ServeHTTP(res, req)
	}
}
