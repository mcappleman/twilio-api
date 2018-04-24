package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gopkg.in/mgo.v2"
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

type stats struct {
	correct		int
	total		int
	percent		float64
}

type buckets struct {
	all				stats
	bucket50_55		stats
	bucket55_60		stats
	bucket60_65		stats
	bucket65_70		stats
	bucket70_100	stats
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

	message, err := getBuckets(bc.DB)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = sendMessage(message, r.Form["From"][0], r.Form["To"][0])
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

func getBuckets(db *mgo.Database) (string, error) {

	var err error
	currentBuckets := buckets{}

	currentBuckets.all, err = fillBucket(db, 0, 101)
	if err != nil {
		return "", err
	}

	currentBuckets.bucket50_55, err = fillBucket(db, 50, 55)
	if err != nil {
		return "", err
	}

	currentBuckets.bucket55_60, err = fillBucket(db, 55, 60)
	if err != nil {
		return "", err
	}

	currentBuckets.bucket60_65, err = fillBucket(db, 60, 65)
	if err != nil {
		return "", err
	}

	currentBuckets.bucket65_70, err = fillBucket(db, 65, 70)
	if err != nil {
		return "", err
	}

	currentBuckets.bucket70_100, err = fillBucket(db, 70, 101)
	if err != nil {
		return "", err
	}

	returnString := fmt.Sprintf("\nAll: %d of %d for %f percent accuracy", currentBuckets.all.correct, currentBuckets.all.total, currentBuckets.all.percent)
	returnString += fmt.Sprintf("\n50-55: %d of %d for %f percent accuracy", currentBuckets.bucket50_55.correct, currentBuckets.bucket50_55.total, currentBuckets.bucket50_55.percent)
	returnString += fmt.Sprintf("\n55-60: %d of %d for %f percent accuracy", currentBuckets.bucket55_60.correct, currentBuckets.bucket55_60.total, currentBuckets.bucket55_60.percent)
	returnString += fmt.Sprintf("\n60-65: %d of %d for %f percent accuracy", currentBuckets.bucket60_65.correct, currentBuckets.bucket60_65.total, currentBuckets.bucket60_65.percent)
	returnString += fmt.Sprintf("\n65-70: %d of %d for %f percent accuracy", currentBuckets.bucket65_70.correct, currentBuckets.bucket65_70.total, currentBuckets.bucket65_70.percent)
	returnString += fmt.Sprintf("\n70-100: %d of %d for %f percent accuracy", currentBuckets.bucket70_100.correct, currentBuckets.bucket70_100.total, currentBuckets.bucket70_100.percent)

	return returnString, nil

}

func fillBucket(db *mgo.Database, min int, max int) (stats, error) {

	gList, err := models.GetBucket(db, float64(min), float64(max))
	if err != nil {
		return stats{}, err
	}

	currentStats := stats{}
	currentStats.total = len(gList)
	var correctPicks int = 0

	for _, game := range gList {

		winner := game.HomeTeam
		if game.AwayRuns > game.HomeRuns {
			winner = game.AwayTeam
		}

		if winner == game.NumberFireFavorite {
			correctPicks += 1
		}

	}

	currentStats.correct = correctPicks
	currentStats.percent = float64(correctPicks)/float64(len(gList))

	return currentStats, nil

}
