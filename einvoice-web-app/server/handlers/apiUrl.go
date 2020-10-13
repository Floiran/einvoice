package handlers

import (
	. "github.com/slovak-egov/einvoice/einvoice-web-app/server/config"
	"net/http"
)

func ApiUrlHandler(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(Config.AuthServerUrl))
}
