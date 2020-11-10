package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/slovak-egov/einvoice/apiserver/config"
	"github.com/slovak-egov/einvoice/apiserver/manager"
	"github.com/slovak-egov/einvoice/apiserver/xml"
	"github.com/slovak-egov/einvoice/handlers"
)

type App struct {
	router    *mux.Router
	config    config.Configuration
	Manager   manager.Manager
	validator xml.Validator
}

func NewApp() *App {
	appConfig := config.Init()

	a := &App{
		config: appConfig,
		Manager: manager.Init(appConfig),
		validator: xml.NewValidator(appConfig),
		router: mux.NewRouter(),
	}

	a.InitializeHandlers()

	return a
}

func (a *App) InitializeHandlers() {
	a.router.HandleFunc("/invoices", a.getInvoices).Methods("GET")
	a.router.HandleFunc("/invoices", a.createInvoice).Methods("POST")
	// Maybe we can merge following 2 endpoints
	a.router.HandleFunc("/invoices/{id:[0-9]+}", a.getInvoice).Methods("GET")
	a.router.HandleFunc("/invoices/{id:[0-9]+}/detail", a.getInvoiceDetail).Methods("GET")
	a.router.HandleFunc("/attachments/{id:[0-9]+}", a.getAttachment).Methods("GET")
}

func (a *App) Run() {
	srv := &http.Server{
		Handler:      handlers.LoggingHandler{a.router},
		Addr:         fmt.Sprintf("%s:%d", "0.0.0.0", a.config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.WithField("address", srv.Addr).Info("app.server_start")

	log.Fatal(srv.ListenAndServe())
}

func (a *App) Close() {
	// TODO: https://github.com/gorilla/mux#graceful-shutdown
	a.Manager.Db.Close()
}

func (a *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	a.router.ServeHTTP(w, req)
}
