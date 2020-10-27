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
	authenticator := auth.NewAuthenticator(userManager)

	keys := auth.NewKeys()

	router := mux.NewRouter()

	router.PathPrefix("/login").HandlerFunc(auth.HandleLogin(userManager, keys))
	router.PathPrefix("/logout").HandlerFunc(authenticator.WithUser(auth.HandleLogout(userManager)))
	// TODO: once user table is moved to postgres change url to /users/:id
	// Check if current user has access to user:id data
	router.PathPrefix("/users/me").Methods("GET").HandlerFunc(authenticator.WithUser(auth.HandleMe))
	router.PathPrefix("/users/me").Methods("PUT").HandlerFunc(authenticator.WithUser(auth.HandleUpdateUser(userManager)))

	proxy := httputil.NewSingleHostReverseProxy(apiserver)

	router.PathPrefix("/invoices").Methods("GET").HandlerFunc(auth.HandleOpenProxy(proxy))
	router.PathPrefix("/attachments").Methods("GET").HandlerFunc(auth.HandleOpenProxy(proxy))

	router.PathPrefix("/").HandlerFunc(authenticator.WithUser(auth.HandleAuthProxy(proxy)))

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
	handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"}),
}
