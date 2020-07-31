package main

import (
	"github.com/filipsladek/einvoice/common"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	var apiServerUrl = common.GetRequiredEnvVariable("API_SERVER_URL")

	var clientBuildDir = "../client/build/"
	var entry = clientBuildDir + "/index.html"

	var port = "8081"

	rand.Seed(time.Now().Unix())

	r := mux.NewRouter()

	r.Path("/").HandlerFunc(IndexHandler(entry))

	r.Path("/api/url").HandlerFunc(ApiUrlHandler(apiServerUrl))

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

func ApiUrlHandler(apiServerUrl string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(apiServerUrl))
	}

	return http.HandlerFunc(fn)
}
