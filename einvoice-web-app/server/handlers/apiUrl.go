package handlers

import (
	"github.com/slovak-egov/einvoice/common"
	"net/http"
)

func ApiUrlHandler(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(common.GetRequiredEnvVariable("API_SERVER_URL")))
}
