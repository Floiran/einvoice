package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/slovak-egov/einvoice/common"
	apiHandlers "github.com/slovak-egov/einvoice/einvoice-web-app/server/handlers"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	clientBuildDir := common.GetRequiredEnvVariable("CLIENT_BUILD_DIR")

	port := common.GetRequiredEnvVariable("PORT")

	rand.Seed(time.Now().Unix())

	router := mux.NewRouter()

	router.Path("/api/url").HandlerFunc(apiHandlers.ApiUrlHandler)

	router.PathPrefix("/").Handler(
		apiHandlers.UiHandler{StaticPath: clientBuildDir, IndexPath: "index.html"},
	)

	fmt.Println("server start")

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, router),
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
