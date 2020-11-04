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

	user := &db.User{}

	if err := json.Unmarshal(body, &user); err != nil {
		handlers.RespondWithError(res, http.StatusBadRequest, "Invalid payload")
		return
	}
	user.Id = req.Header.Get("User-Id")

	user, err = a.manager.UpdateUser(user)
	if err != nil {
		handlers.RespondWithError(res, http.StatusInternalServerError, "Something went wrong")
		return
	}

	handlers.RespondWithJSON(res, http.StatusOK, user)
}
