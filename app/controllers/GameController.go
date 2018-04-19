package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/mcappleman/twilio-api/app/models"
)

type ListResponse struct {
	Message string
	Body    []models.Game
}

type TwilioParams struct {
	MessageSid          string
	SmsSid              string
	AccountSid          string
	MessagingServiceSid string
	From                string
	To                  string
	Body                string
	NumMedia            string
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

	/*for key, value := range r.Form {
		fmt.Printf("%s = %s\n", key, value)
	}*/

	err := sendMessage("Hello Twilio World", r.Form["From"][0], r.Form["To"][0])
	if err != nil {
		fmt.Println(err)
		return
	}

}

func sendMessage(message string, toNumber string, fromNumber string) error {

	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	twilioUrl := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	msgData := url.Values{}
	msgData.Set("To", toNumber)
	msgData.Set("From", fromNumber)
	msgData.Set("Body", message)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, err := http.NewRequest("POST", twilioUrl, &msgDataReader)
	if err != nil {
		return err
	}

	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {

		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err != nil {
			return err
		}

	}

	return nil

}
