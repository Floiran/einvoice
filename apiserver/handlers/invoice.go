package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/slovak-egov/einvoice/apiserver/attachment"
	"github.com/slovak-egov/einvoice/apiserver/manager"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetInvoiceMetaHandler(manager manager.Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		invoiceId, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(400)
			return
		}

		meta, err := manager.GetMeta(invoiceId)
		if err != nil {
			w.WriteHeader(404)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(meta)
	}
}

func GetFullInvoiceHandler(manager manager.Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(400)
			return
		}

		inv, err := manager.GetFull(id)
		if err != nil {
			w.WriteHeader(404)
			return
		}

		w.Header().Set("Content-Type", "application/xml")
		w.Write([]byte(inv))
	}
}

func GetAttachmentHandler(manager manager.Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(400)
			return
		}

		meta, err := manager.GetAttachmentMeta(id)
		if err != nil {
			w.WriteHeader(404)
			return
		}

		at, err := manager.GetAttachmentFile(id)
		if err != nil {
			w.WriteHeader(404)
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename="+meta.Name)
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		w.Write([]byte(at))
	}
}

func GetAllInvoicesHandler(manager manager.Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		invoices, err := manager.GetAllInvoiceMeta()
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if len(invoices) == 0 {
			w.Write([]byte("[]"))
		} else {
			json.NewEncoder(w).Encode(invoices)
		}
	}
}

func CreateInvoiceHandler(manager manager.Manager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		data := r.PostFormValue("data")
		format := r.PostFormValue("format")
		ats, err := getAttachments(r)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		meta, err := manager.CreateInvoice(format, data, ats)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(meta); err != nil {
			panic(err)
		}
	}
}

func getAttachments(r *http.Request) ([]*attachment.PostAttachment, error) {
	var ats []*attachment.PostAttachment
	for k, _ := range r.MultipartForm.File {
		if strings.HasPrefix(k, "attachment") {
			file, handler, err := r.FormFile(k)
			if err != nil {
				return nil, err
			}
			defer file.Close()

			bytes, err := ioutil.ReadAll(file)
			if err != nil {
				return nil, err
			}

			ats = append(ats,
				&attachment.PostAttachment{
					Name:    handler.Filename,
					Content: bytes,
				})
		}
	}
	return ats, nil
}
