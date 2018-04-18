package controllers

import (
	// "fmt"
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2"
)

type BaseController struct {
	DB	*mgo.Database
}

func Init(db *mgo.Database) BaseController {

	bc := BaseController{}
	bc.DB = db
	return bc

}

func (bc *BaseController) RespondWithError(w http.ResponseWriter, code int, message string) {

	bc.RespondWithJson(w, code, map[string]string{"error": message})

}

func (bc *BaseController) RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {

	// fmt.Println(len(payload))
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}

func (bc *BaseController) Index(w http.ResponseWriter, r *http.Request) {

	bc.RespondWithJson(w, http.StatusOK, map[string]string{"message": "Hello Twilio World!"})

}
