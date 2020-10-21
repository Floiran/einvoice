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

func handleRequests(manager manager.Manager) {
	router := mux.NewRouter()

	router.Path("/api/invoices").Methods("GET").HandlerFunc(apiHandlers.GetAllInvoicesHandler(manager))
	router.Path("/api/invoices").Methods("POST").HandlerFunc(apiHandlers.CreateInvoiceHandler(manager))
	router.Path("/api/invoices/{id}").Methods("GET").HandlerFunc(apiHandlers.GetInvoiceMetaHandler(manager))
	router.Path("/api/invoices/{id}/full").Methods("GET").HandlerFunc(apiHandlers.GetFullInvoiceHandler(manager))
	router.Path("/api/attachments/{id}").Methods("GET").HandlerFunc(apiHandlers.GetAttachmentHandler(manager))

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

	manager := manager.NewManager(db, storage, validator)

	handleRequests(manager)
}
