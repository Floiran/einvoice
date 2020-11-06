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
	Router    *mux.Router
	Config    config.Configuration
	Manager   manager.Manager
	validator xml.Validator
}

func (a *App) Initialize() {
	a.Config = config.Init()
	a.validator = xml.NewValidator(a.Config)
	a.Manager = manager.Init(a.Config)

	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/invoices", a.getInvoices).Methods("GET")
	a.Router.HandleFunc("/invoices", a.createInvoice).Methods("POST")
	// Maybe we can merge following 2 endpoints
	a.Router.HandleFunc("/invoices/{id:[0-9]+}", a.getInvoice).Methods("GET")
	a.Router.HandleFunc("/invoices/{id:[0-9]+}/detail", a.getInvoiceDetail).Methods("GET")
	a.Router.HandleFunc("/attachments/{id:[0-9]+}", a.getAttachment).Methods("GET")
}

func (a *App) Run() {
	srv := &http.Server{
		Handler:      handlers.LoggingHandler{a.Router},
		Addr:         fmt.Sprintf("%s:%d", "0.0.0.0", a.Config.Port),
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
