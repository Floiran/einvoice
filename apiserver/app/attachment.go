package app

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/slovak-egov/einvoice/handlers"
)

func (a *App) getAttachment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handlers.RespondWithError(w, http.StatusBadRequest, "ID should be integer")
		return
	}

	attachment, name, err := a.manager.GetAttachment(id)
	if err != nil {
		handlers.RespondWithError(w, http.StatusNotFound, "Attachment was not found")
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+name)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Write([]byte(attachment))
}
