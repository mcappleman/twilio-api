package controllers

import (
	"net/http"

	"github.com/mcappleman/twilio-api/app/models"
	"github.com/gorilla/mux"
)

func (a *App) GetGames(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	gList, err := models.GetGames(a.DB.Database())
	if err != nil {
		a.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.ReturnWithJson(w, http.StatusOK, gList)

}
