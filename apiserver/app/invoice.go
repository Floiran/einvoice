package app

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/slovak-egov/einvoice/apiserver/db"
	"github.com/slovak-egov/einvoice/apiserver/manager"
)

func (a *App) getInvoices(w http.ResponseWriter, r *http.Request) {
	formats := r.URL.Query()["format"]
	if len(formats) == 0 {
		formats = []string{db.UblFormat, db.D16bFormat}
	}

	invoices, err := a.manager.GetInvoices(formats)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respondWithJSON(w, http.StatusOK, invoices)
}

func (a *App) getInvoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID should be integer")
		return
	}

	invoice, err := a.manager.GetInvoice(id)
	if err != nil {
		// TODO: distinguish NotFound and other errors
		respondWithError(w, http.StatusNotFound, "Invoice was not found")
		return
	}

	respondWithJSON(w, http.StatusOK, invoice)
}

func (a *App) getInvoiceDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID should be integer")
		return
	}

	invoice, err := a.manager.GetInvoiceDetail(id)
	if err != nil {
		// TODO: distinguish NotFound and other errors
		respondWithError(w, http.StatusNotFound, "Invoice was not found")
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(invoice))
}

func (a *App) createInvoice(w http.ResponseWriter, r *http.Request) {
	// TODO: return 413 if request is too large
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	data := r.PostFormValue("data")
	format := r.PostFormValue("format")
	ats, err := parseAttachments(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid attachments")
		return
	}

	// Validate payload
	switch format {
	case db.UblFormat:
		if err = a.validator.ValidateUBL21([]byte(data)); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invoice is not valid")
			return
		}
	case db.D16bFormat:
		if err = a.validator.ValidateD16B([]byte(data)); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invoice is not valid")
			return
		}
	default:
		respondWithError(w, http.StatusBadRequest, "Invalid format")
		return
	}

	metadata, err := a.manager.CreateInvoice(format, data, ats)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respondWithJSON(w, http.StatusCreated, metadata)
}

func parseAttachments(r *http.Request) ([]*manager.Attachment, error) {
	var ats []*manager.Attachment
	for k := range r.MultipartForm.File {
		if strings.HasPrefix(k, "attachment") {
			file, handler, err := r.FormFile(k)
			if err != nil {
				return nil, err
			}

			bytes, err := ioutil.ReadAll(file)
			if err != nil {
				return nil, err
			}
			file.Close()

			ats = append(
				ats,
				&manager.Attachment{
					Name:    handler.Filename,
					Content: bytes,
				},
			)
		}
	}
	return ats, nil
}
