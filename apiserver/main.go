// main.go
package main

import (
	"fmt"
	apiHandlers "github.com/filipsladek/einvoice/apiserver/handlers"
	"github.com/filipsladek/einvoice/apiserver/invoice"
	"github.com/filipsladek/einvoice/apiserver/manager"
	"github.com/filipsladek/einvoice/apiserver/postgres"
	"github.com/filipsladek/einvoice/apiserver/storage"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

func handleRequests(manager manager.Manager) {
	router := mux.NewRouter()

	router.PathPrefix("/api/invoices").Methods("GET").HandlerFunc(apiHandlers.GetAllInvoicesHandler(manager))
	router.PathPrefix("/api/invoice/full/{id}").Methods("GET").HandlerFunc(apiHandlers.GetFullInvoiceHandler(manager))
	router.PathPrefix("/api/invoice/meta/{id}").Methods("GET").HandlerFunc(apiHandlers.GetInvoiceMetaHandler(manager))
	router.PathPrefix("/api/invoice/json").Methods("POST").HandlerFunc(apiHandlers.CreateInvoiceJsonHandler(manager))
	router.PathPrefix("/api/invoice/ubl").Methods("POST").HandlerFunc(apiHandlers.CreateInvoiceXmlUblHandler(manager))
	router.PathPrefix("/api/invoice/d16b").Methods("POST").HandlerFunc(apiHandlers.CreateInvoiceXmlD16bHandler(manager))

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, router),
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	println("Server running on", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}

func main() {
	fmt.Println("start")
	storage := storage.InitStorage()
	storage.SaveObject("abc", "def")
	fmt.Println("stored")

	dbConf := postgres.NewConnectionConfig()

	db := postgres.Connect(dbConf)
	defer db.Close()

	dbConnector := &postgres.DBConnector{DB: db}

	postgres.InitDB(dbConnector)

	manager := manager.NewManager(dbConnector, storage)

	// dummy data
	if len(dbConnector.GetAllInvoice()) == 0 {
		manager.Create(&invoice.Invoice{Sender: "SubjectA", Receiver: "SubjectB", Price: 100})
		manager.Create(&invoice.Invoice{Sender: "SubjectA", Receiver: "SubjectC", Price: 200})
		manager.Create(&invoice.Invoice{Sender: "SubjectA", Receiver: "SubjectD", Price: 300})
	}

	handleRequests(manager)
}
