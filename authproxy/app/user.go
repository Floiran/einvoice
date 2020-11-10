package app

import (
	"encoding/json"
	"net/http"

	"github.com/slovak-egov/einvoice/authproxy/db"
	"github.com/slovak-egov/einvoice/handlers"
)

func (a *App) getUser(res http.ResponseWriter, req *http.Request) {
	userId := req.Header.Get("User-Id")
	user, err := a.manager.GetUser(userId)
	if err != nil {
		handlers.RespondWithError(res, http.StatusInternalServerError, "Something went wrong")
		return
	}
	handlers.RespondWithJSON(res, http.StatusOK, user)
}

func (a *App) updateUser(res http.ResponseWriter, req *http.Request) {
	var requestBody *db.UserUpdate

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&requestBody); err != nil {
		handlers.RespondWithError(res, http.StatusBadRequest, err.Error())
		return
	}

	user, err := a.manager.UpdateUser(req.Header.Get("User-Id"), requestBody)
	if err != nil {
		handlers.RespondWithError(res, http.StatusInternalServerError, "Something went wrong")
		return
	}

	handlers.RespondWithJSON(res, http.StatusOK, user)
}
