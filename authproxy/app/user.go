package app

import (
	"encoding/json"
	"io"
	"io/ioutil"
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
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		handlers.RespondWithError(res, http.StatusBadRequest, "Invalid payload")
		return
	}
	if err = req.Body.Close(); err != nil {
		handlers.RespondWithError(res, http.StatusBadRequest, "Invalid payload")
		return
	}

	updatedUserData := &db.UserUpdate{}

	if err := json.Unmarshal(body, &updatedUserData); err != nil {
		handlers.RespondWithError(res, http.StatusBadRequest, "Invalid payload")
		return
	}
	if updatedUserData.IsEmpty() {
		handlers.RespondWithError(res, http.StatusBadRequest, "Empty body")
		return
	}

	updatedUserData.UserId = req.Header.Get("User-Id")

	user, err := a.manager.UpdateUser(updatedUserData)
	if err != nil {
		handlers.RespondWithError(res, http.StatusInternalServerError, "Something went wrong")
		return
	}

	handlers.RespondWithJSON(res, http.StatusOK, user)
}
