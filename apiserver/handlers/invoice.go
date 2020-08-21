package handlers

import (
	"encoding/json"
	"encoding/xml"
	. "github.com/filipsladek/einvoice/apiserver/invoice"
	"github.com/filipsladek/einvoice/apiserver/postgres"
	"github.com/filipsladek/einvoice/apiserver/storage"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

func GetInvoiceMetaHandler(storage storage.Storage, db postgres.DBConnector) func(w http.ResponseWriter, r *http.Request) {
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

		w.Header().Set("Content-Type", "application/xml")
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

func CreateInvoiceJsonHandler(storage storage.Storage, db postgres.DBConnector) func(w http.ResponseWriter, r *http.Request) {
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

		meta := invoice.GetMeta()
		db.CreateInvoice(meta)
		invoice.Id = meta.Id

		xmlString, err := xml.Marshal(invoice)
		if err != nil {
			panic(err)
		}
		err = storage.SaveObject("invoice-"+invoice.Id, string(xmlString))
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(meta); err != nil {
			panic(err)
		}
	}
}

func CreateInvoiceXmlHandler(storage storage.Storage, db postgres.DBConnector) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var invoice = new(Invoice)
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}

		if err := xml.Unmarshal(body, &invoice); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

		meta := invoice.GetMeta()
		db.CreateInvoice(meta)
		invoice.Id = meta.Id

		xmlString, err := xml.Marshal(invoice)
		if err != nil {
			panic(err)
		}
		err = storage.SaveObject("invoice-"+invoice.Id, string(xmlString))
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(meta); err != nil {
			panic(err)
		}
	}
}
