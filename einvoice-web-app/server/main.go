package main

import (
	"encoding/json"
	"github.com/filipsladek/einvoice/common"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	var clientBuildDir = "../client/build"
	var entry = clientBuildDir + "/index.html"

	var port = "8081"

	rand.Seed(time.Now().Unix())

	r := mux.NewRouter()

	r.Path("/").HandlerFunc(IndexHandler(entry))

	r.PathPrefix("/invoice/{id}").Methods("GET").HandlerFunc(GetInvoiceHandler)
	r.PathPrefix("/invoice").Methods("POST").HandlerFunc(CreateInvoiceHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(clientBuildDir)))

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, r),
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}

func GetInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invoiceId := vars["id"]

	invoice := common.NewInvoice(invoiceId, "subject"+strconv.Itoa(rand.Intn(100)), "subject"+strconv.Itoa(rand.Intn(100)))

	json.NewEncoder(w).Encode(invoice)
}

func CreateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	var invoice common.Invoice
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	invoice.Id = strconv.Itoa(rand.Intn(100000))

	if err := json.Unmarshal(body, &invoice); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(invoice); err != nil {
		panic(err)
	}
}
