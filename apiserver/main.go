// main.go
package main

import (
	"encoding/json"
	"fmt"
	apiHandlers "github.com/filipsladek/einvoice/apiserver/handlers"
	"github.com/filipsladek/einvoice/apiserver/invoice"
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

func handleRequests(storage storage.Storage, db postgres.DBConnector) {
	router := mux.NewRouter()

	router.PathPrefix("/api/invoice/full/{id}").Methods("GET").HandlerFunc(apiHandlers.GetFullInvoiceHandler(storage, db))
	router.PathPrefix("/api/invoice/{id}").Methods("GET").HandlerFunc(apiHandlers.GetInvoiceHandler(storage, db))
	router.PathPrefix("/api/invoices").Methods("GET").HandlerFunc(apiHandlers.GetAllInvoicesHandler(storage, db))
	router.PathPrefix("/api/invoice").Methods("POST").HandlerFunc(apiHandlers.CreateInvoiceHandler(storage, db))

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, handlers.CORS(corsOptions...)(router)),
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func createDummyInvoice(invoice invoice.Invoice, dbConnector postgres.DBConnector, storage storage.Storage) {
	dbConnector.CreateInvoice(&invoice)
	json, _ := json.Marshal(invoice)
	storage.SaveObject("invoice-"+invoice.Id, string(json))
}

func main() {
	fmt.Println("start")
	storage := storage.InitStorage()
	storage.SaveObject("abc", "def")
	fmt.Println("stored")

	dbConf := postgres.NewConnectionConfig()

	db := postgres.Connect(dbConf)
	defer db.Close()

	dbConnector := postgres.DBConnector{DB: db}

	postgres.InitDB(dbConnector)

	// dummy data
	if len(dbConnector.GetAllInvoice()) == 0 {
		createDummyInvoice(invoice.Invoice{Sender: "SubjectA", Receiver: "SubjectB"}, dbConnector, storage)
		createDummyInvoice(invoice.Invoice{Sender: "SubjectA", Receiver: "SubjectC"}, dbConnector, storage)
		createDummyInvoice(invoice.Invoice{Sender: "SubjectA", Receiver: "SubjectD"}, dbConnector, storage)
	}

	handleRequests(storage, dbConnector)
}

var corsOptions = []handlers.CORSOption{
	handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Accept", "token"}),
	handlers.AllowedOrigins([]string{"*"}),
	handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
}
