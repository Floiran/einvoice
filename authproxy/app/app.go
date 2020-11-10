package app

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	muxHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/slovak-egov/einvoice/authproxy/config"
	"github.com/slovak-egov/einvoice/authproxy/manager"
	"github.com/slovak-egov/einvoice/authproxy/slovenskoSk"
	"github.com/slovak-egov/einvoice/handlers"
)

type App struct {
	router      *mux.Router
	config      config.Configuration
	manager     manager.Manager
	slovenskoSk slovenskoSk.Connector
}

func NewApp() *App {
	appConfig := config.Init()
	a := &App{
		config:  appConfig,
		manager: manager.Init(appConfig),
		slovenskoSk: slovenskoSk.Init(appConfig.SlovenskoSk),
	}

	a.InitializeRouter()

	return a
}

func (a *App) InitializeRouter() {
	apiserver, err := url.Parse(a.config.ApiServerUrl)
	if err != nil {
		log.WithField("error", err.Error()).Fatal("app.initialization.apiserver_url.parse")
	}

	a.router = mux.NewRouter()
	authRouter := a.router.PathPrefix("/").Subrouter()
	authRouter.Use(a.authMiddleware)

	a.router.HandleFunc("/login", a.handleLogin).Methods("GET")
	a.router.HandleFunc("/logout", a.handleLogout).Methods("GET")
	// TODO: Change url to /users/:id
	// Check if current user has access to user:id data
	authRouter.HandleFunc("/users/me", a.getUser).Methods("GET")
	authRouter.HandleFunc("/users/me", a.updateUser).Methods("PATCH")

	proxy := httputil.NewSingleHostReverseProxy(apiserver)

	a.router.Methods("GET").HandlerFunc(handleProxy(proxy))

	authRouter.Methods("POST").HandlerFunc(handleProxy(proxy))
}

func (a *App) Run() {
	srv := &http.Server{
		Handler:      handlers.LoggingHandler{muxHandlers.CORS(corsOptions...)(a.router)},
		Addr:         fmt.Sprintf("%s:%d", "0.0.0.0", a.config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.WithField("address", srv.Addr).Info("app.server_start")

	log.Fatal(srv.ListenAndServe())
}

func (a *App) Close() {
	// TODO: https://github.com/gorilla/mux#graceful-shutdown
	a.manager.Db.Close()
}

func (a *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	a.router.ServeHTTP(w, req)
}

var corsOptions = []muxHandlers.CORSOption{
	muxHandlers.AllowedHeaders([]string{"Content-Type", "Origin", "Accept", "token", "Authorization"}),
	muxHandlers.AllowedOrigins([]string{"*"}),
	muxHandlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH"}),
}

func handleProxy(proxy *httputil.ReverseProxy) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host
		proxy.ServeHTTP(res, req)
	}
}
