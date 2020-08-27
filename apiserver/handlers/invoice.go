package handlers

import (
	"encoding/json"
	"github.com/filipsladek/einvoice/apiserver/manager"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

func GetInvoiceMetaHandler(manager manager.Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		invoiceId := vars["id"]

		err, meta := manager.GetMeta(invoiceId)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(meta)
	}
}

func GetFullInvoiceHandler(manager manager.Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		err, inv := manager.Get(id)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/xml")
		json.NewEncoder(w).Encode(inv)
	}
}

func GetAllInvoicesHandler(manager manager.Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		invoices := manager.GetAllInvoiceMeta()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(invoices)
	}
}

func CreateInvoiceJsonHandler(manager manager.Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}

		err, _, meta := manager.CreateJSON(string(body))
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

//func CreateInvoiceXmlHandler(manager invoice.Manager) func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var invoice = new(Invoice)
//		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
//		if err != nil {
//			panic(err)
//		}
//		if err := r.Body.Close(); err != nil {
//			panic(err)
//		}
//
//		if err := xml.Unmarshal(body, &invoice); err != nil {
//			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//			w.WriteHeader(422)
//			if err := json.NewEncoder(w).Encode(err); err != nil {
//				panic(err)
//			}
//		}
//
//		meta := invoice.GetMeta()
//		db.CreateInvoice(meta)
//		invoice.Id = meta.Id
//
//		xmlString, err := xml.Marshal(invoice)
//		if err != nil {
//			panic(err)
//		}
//		err = storage.SaveObject("invoice-"+invoice.Id, string(xmlString))
//		if err != nil {
//			panic(err)
//		}
//
//		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//		w.WriteHeader(http.StatusCreated)
//		if err := json.NewEncoder(w).Encode(meta); err != nil {
//			panic(err)
//		}
//	}
//}
