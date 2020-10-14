package handlers

import (
	"encoding/json"
	"github.com/slovak-egov/einvoice/einvoice-web-app/server/config"
	"net/http"
)

type urls struct {
	ApiServerUrl        string `json:"apiServerUrl"`
	SlovenskoSkLoginUrl string `json:"slovenskoSkLoginUrl"`
}

func ApiUrlHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(urls{
		config.Config.AuthServerUrl,
		config.Config.SlovenskoSkLoginUrl,
	})
}
