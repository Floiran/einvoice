package app

import (
	"net/http"

	"github.com/slovak-egov/einvoice/authproxy/db"
	"github.com/slovak-egov/einvoice/handlers"
)

func (a *App) handleLogin(res http.ResponseWriter, req *http.Request) {
	oboToken, err := GetBearerToken(req)
	if err != nil {
		handlers.RespondWithError(res, http.StatusUnauthorized, err.Error())
		return
	}
	slovenskoSkUser, err := a.slovenskoSk.GetUser(oboToken)
	if err != nil {
		handlers.RespondWithError(res, http.StatusUnauthorized, "Unauthorized")
		return
	}

	id := slovenskoSkUser.Uri

	// TODO: SELECT or INSERT query
	user, err := a.manager.GetUser(id)
	if user == nil {
		user, err = a.manager.CreateUser(id, slovenskoSkUser.Name)
	}

	if err != nil {
		handlers.RespondWithError(res, http.StatusInternalServerError, "Something went wrong")
		return
	}

	token := a.manager.CreateUserToken(id)

	handlers.RespondWithJSON(res, http.StatusOK, struct{
		Token string `json:"token"`
		*db.User
	}{token, user})
}

func (a *App) handleLogout(res http.ResponseWriter, req *http.Request) {
	token, err := GetBearerToken(req)
	if err != nil {
		handlers.RespondWithError(res, http.StatusUnauthorized, err.Error())
		return
	}

	err = a.manager.LogoutUser(token)
	if err != nil {
		handlers.RespondWithError(res, http.StatusUnauthorized, "Unauthorized")
		return
	}
	handlers.RespondWithJSON(res, http.StatusOK, map[string]string{"logout": "successful"})
}
