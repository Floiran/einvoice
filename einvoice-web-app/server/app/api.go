package app

import (
	"encoding/json"
	"net/http"
)

func (a *App) ApiUrlHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"apiServerUrl": a.Config.AuthServerUrl,
		"slovenskoSkLoginUrl": a.Config.SlovenskoSkLoginUrl,
	})
}
