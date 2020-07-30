package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var clientBuildDir = "../client/build"

	var entry = clientBuildDir + "/index.html"
	var port = "8081"

	r := mux.NewRouter()

	r.Path("/").HandlerFunc(IndexHandler(entry))

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

	bytes, _ := ioutil.ReadFile(entrypoint)
	println("str", string(bytes))

	return http.HandlerFunc(fn)
}
