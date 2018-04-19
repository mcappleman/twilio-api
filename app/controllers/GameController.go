package controllers

import (
	"fmt"
	// "encoding/json"
	"net/http"

	"github.com/mcappleman/twilio-api/app/models"
)

type ListResponse struct {
	Message	string
	Body	[]models.Game
}

type TwilioParams struct {
	MessageSid			string
	SmsSid				string
	AccountSid			string
	MessagingServiceSid	string
	From				string
	To					string
	Body				string
	NumMedia			string
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

func (bc *BaseController) PostMessage(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	for key, value := range r.Form {
		fmt.Printf("%s = %s\n", key, value)
	}

	fmt.Println("Message Successful")

}
