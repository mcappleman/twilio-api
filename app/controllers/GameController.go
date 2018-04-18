package controllers

import (
	// "fmt"
	"net/http"

	"github.com/mcappleman/twilio-api/app/models"
)

type ListResponse struct {
	Message	string
	Body	[]models.Game
}

func (bc *BaseController) GetGames(w http.ResponseWriter, r *http.Request) {

	// vars := mux.Vars(r)
	gList, err := models.GetGames(bc.DB)
	if err != nil {
		bc.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// fmt.Println(len(gList))
	payload := ListResponse{"Games Retrieved Successfully", gList}

	bc.RespondWithJson(w, http.StatusOK, payload)

}
