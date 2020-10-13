package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/slovak-egov/einvoice/authproxy/auth"
	. "github.com/slovak-egov/einvoice/authproxy/config"
	"github.com/slovak-egov/einvoice/authproxy/db"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

func main() {
	InitConfig()

	apiserver, err := url.Parse(Config.ApiServerUrl)
	if err != nil {
		panic(err)
	}

	authDB := db.NewAuthDB()
	userManager := auth.NewUserManager(authDB)

	router := mux.NewRouter()

	router.PathPrefix("/login").HandlerFunc(auth.HandleLogin(userManager))
	router.PathPrefix("/logout").HandlerFunc(auth.HandleLogout(userManager))
	router.PathPrefix("/me").HandlerFunc(auth.HandleMe(userManager))

	proxy := httputil.NewSingleHostReverseProxy(apiserver)
	router.PathPrefix("/").HandlerFunc(auth.WithToken(userManager, func(res http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host
		proxy.ServeHTTP(res, req)
	}))

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, handlers.CORS(corsOptions...)(router)),
		Addr:         fmt.Sprintf("%s:%d", "0.0.0.0", Config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server running on", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}

var corsOptions = []handlers.CORSOption{
	handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Accept", "token", "Authorization"}),
	handlers.AllowedOrigins([]string{"*"}),
	handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
}
