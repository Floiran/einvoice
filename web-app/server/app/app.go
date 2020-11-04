package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/slovak-egov/einvoice/handlers"
	"github.com/slovak-egov/einvoice/web-app/server/config"
)

type App struct {
	Router *mux.Router
	Config config.Configuration
}

func (a *App) Initialize() {
	a.Config = config.Init()

	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/api/urls", a.ApiUrlHandler).Methods("Get")

	a.Router.PathPrefix("/").Handler(
		UiHandler{StaticPath: a.Config.ClientBuildDir, IndexPath: "index.html"},
	)
}

func (a *App) Run() {
	srv := &http.Server{
		Handler:      handlers.LoggingHandler{a.Router},
		Addr:         fmt.Sprintf("%s:%d", "0.0.0.0", a.Config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.WithField("address", srv.Addr).Info("app.server_start")

	log.Fatal(srv.ListenAndServe())
}
