package app

import (
	"net/http"

	"github.com/slovak-egov/einvoice/handlers"
)

func (a *App) handleLogin(res http.ResponseWriter, req *http.Request) {
	slovenskoSkUser, err := a.slovenskoSk.GetUser(req.Header.Get("Authorization"))
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

	handlers.RespondWithJSON(res, http.StatusOK, map[string] string{
		"token": token,
		"id": id,
		"email": user.Email,
		"serviceAccountKey": user.ServiceAccountKey,
		"name": user.Name,
	})
}

func (a *App) handleLogout(res http.ResponseWriter, req *http.Request) {
	err := a.manager.LogoutUser(req.Header.Get("Authorization"))
	if err != nil {
		handlers.RespondWithError(res, http.StatusUnauthorized, "Unauthorized")
		return
	}
	handlers.RespondWithJSON(res, http.StatusOK, map[string]string{"logout": "successful"})
}
