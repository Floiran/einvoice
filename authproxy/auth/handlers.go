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

		usr, err := manager.GetUser(id)
		if usr == nil {
			name := slovenskoSkUser.Name
			usr, err = manager.Create(id, name)
		}

		if err != nil {
			res.WriteHeader(500)
			return
		}

		manager.CreateToken(usr)

		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(usr)
	}
}

func HandleLogout(manager UserManager) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		if manager.RemoveToken(req.Header.Get("Authorization")) {
			res.WriteHeader(200)
		} else {
			res.WriteHeader(401)
		}
	}
}

func HandleMe(manager UserManager) func(res http.ResponseWriter, req *http.Request, userId string) {
	return func(res http.ResponseWriter, req *http.Request, userId string) {
		usr, err := manager.GetUser(userId)
		if err != nil {
			res.WriteHeader(500)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(usr)
	}
}

func HandleUpdateUser(manager UserManager) func(res http.ResponseWriter, req *http.Request, userId string) {
	return func(res http.ResponseWriter, req *http.Request, userId string) {
		body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := req.Body.Close(); err != nil {
			panic(err)
		}

		usr := &user.User{}

		if err := json.Unmarshal(body, &usr); err != nil {
			res.WriteHeader(500)
			return
		}
		usr.Id = userId

		err = manager.UpdateUser(usr)
		if err != nil {
			res.WriteHeader(500)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(usr)
	}
}

func HandleAuthProxy(proxy *httputil.ReverseProxy) func(res http.ResponseWriter, req *http.Request, userId string) {
	return func(res http.ResponseWriter, req *http.Request, userId string) {
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
