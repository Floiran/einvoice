package handlers

import (
	"encoding/json"
	. "github.com/filipsladek/einvoice/apiserver/invoice"
	"github.com/filipsladek/einvoice/apiserver/postgres"
	"github.com/filipsladek/einvoice/apiserver/storage"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

func GetInvoiceHandler(storage storage.Storage, db postgres.DBConnector) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		invoiceId := vars["id"]

		invoice := db.GetInvoice(invoiceId)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(invoice)
	}
}

func GetFullInvoiceHandler(storage storage.Storage, db postgres.DBConnector) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		invoiceId := vars["id"]

		invoice, err := storage.ReadObject("invoice-" + invoiceId)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(invoice))
	}
}

func GetAllInvoicesHandler(storage storage.Storage, db postgres.DBConnector) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		invoices := db.GetAllInvoice()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(invoices)
	}
}

func CreateInvoiceHandler(storage storage.Storage, db postgres.DBConnector) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var invoice = new(Invoice)
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}

		if err := json.Unmarshal(body, &invoice); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

		db.CreateInvoice(invoice)

		jsonString, err := json.Marshal(invoice)
		if err != nil {
			panic(err)
		}
		err = storage.SaveObject("invoice-"+invoice.Id, string(jsonString))
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(invoice); err != nil {
			panic(err)
		}
	}
}
