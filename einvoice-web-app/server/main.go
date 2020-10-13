package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	. "github.com/slovak-egov/einvoice/einvoice-web-app/server/config"
	apiHandlers "github.com/slovak-egov/einvoice/einvoice-web-app/server/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	InitConfig()

	router := mux.NewRouter()

	router.Path("/api/url").HandlerFunc(apiHandlers.ApiUrlHandler)

	router.PathPrefix("/").Handler(
		apiHandlers.UiHandler{StaticPath: Config.ClientBuildDir, IndexPath: "index.html"},
	)

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, router),
		Addr:         fmt.Sprintf("%s:%d", "0.0.0.0", Config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server running on", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}
