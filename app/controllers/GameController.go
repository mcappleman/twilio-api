package controllers

import (
	"net/http"

	"github.com/mcappleman/twilio-api/app/models"
)

func (bc *BaseController) GetGames(w http.ResponseWriter, r *http.Request) {

	// vars := mux.Vars(r)
	gList, err := models.GetGames(bc.DB)
	if err != nil {
		bc.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	bc.RespondWithJson(w, http.StatusOK, gList)

}
