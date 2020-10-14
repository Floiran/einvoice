package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	. "github.com/slovak-egov/einvoice/apiserver/config"
	"github.com/slovak-egov/einvoice/apiserver/db"
	apiHandlers "github.com/slovak-egov/einvoice/apiserver/handlers"
	"github.com/slovak-egov/einvoice/apiserver/manager"
	"github.com/slovak-egov/einvoice/apiserver/storage"
	"github.com/slovak-egov/einvoice/apiserver/xml"
)

func handleRequests(manager manager.Manager, validator xml.Validator) {
	router := mux.NewRouter()

	// TODO: update URLs to follow REST conventions
	router.PathPrefix("/api/invoices").Methods("GET").HandlerFunc(apiHandlers.GetAllInvoicesHandler(manager))
	router.PathPrefix("/api/invoice/full/{id}").Methods("GET").HandlerFunc(apiHandlers.GetFullInvoiceHandler(manager))
	router.PathPrefix("/api/invoice/meta/{id}").Methods("GET").HandlerFunc(apiHandlers.GetInvoiceMetaHandler(manager))
	router.PathPrefix("/api/invoice/json").Methods("POST").HandlerFunc(apiHandlers.CreateInvoiceJsonHandler(manager))
	router.PathPrefix("/api/invoice/ubl").Methods("POST").HandlerFunc(apiHandlers.CreateInvoiceXmlUblHandler(manager, validator))
	router.PathPrefix("/api/invoice/d16b").Methods("POST").HandlerFunc(apiHandlers.CreateInvoiceXmlD16bHandler(manager, validator))

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, router),
		Addr:         fmt.Sprintf("%s:%d", "0.0.0.0", Config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server running on", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}

func main() {
	InitConfig()

	storage := storage.InitStorage()

	db := db.NewDBConnector()
	db.Connect()
	defer db.Close()

	validator := xml.NewValidator(
		Config.D16bXsdPath,
		Config.Ubl21XsdPath,
	)

	manager := manager.NewManager(db, storage)

	handleRequests(manager, validator)
}
