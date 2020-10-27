package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/slovak-egov/einvoice/apiserver/config"
	"github.com/slovak-egov/einvoice/apiserver/manager"
	"github.com/slovak-egov/einvoice/apiserver/xml"
)

type App struct {
	Router *mux.Router
	Config config.Configuration
	manager manager.Manager
	validator xml.Validator
}

func (a *App) Initialize() {
	a.validator = xml.NewValidator(a.Config)
	a.manager = manager.Init(a.Config)

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
		Handler:      handlers.LoggingHandler(os.Stdout, a.Router),
		Addr:         fmt.Sprintf("%s:%d", "0.0.0.0", a.Config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server running on", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}

func (a *App) Close() {
	// TODO: https://github.com/gorilla/mux#graceful-shutdown
	a.manager.Db.Close()
}
