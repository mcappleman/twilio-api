package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mcappleman/twilio-api/mongodb"
)

var bc BaseController

func initTest() {

	db, err := mongodb.NewSession(os.Getenv("DATABASE_URL"), os.Getenv("DATABASE_NAME"))
	if err != nil {
		panic(err)
	}

	bc = Init(db.Database())

}

func TestIndex(t *testing.T) {

	fmt.Println("Base Controller Index Test Started")
	initTest()

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	bc.Index(w, r)

	if w.Code != 200 {
		t.Fail()
	}

	var body map[string]string
	json.Unmarshal(w.Body.Bytes(), &body)

	if body["message"] != "Hello Twilio World!" {
		t.Fail()
	}

	fmt.Println("Base Controller Index Test Success")

}

func TestRespondWithJson(t *testing.T) {

	fmt.Println("Base Controller RespondWithJson Test Started")

	w := httptest.NewRecorder()

	bc.RespondWithJson(w, 200, map[string]string{"message": "Test"})

	if w.Code != 200 {
		t.Fail()
	}

	var body map[string]string
	json.Unmarshal(w.Body.Bytes(), &body)

	if body["message"] != "Test" {
		t.Fail()
	}

	fmt.Println("Base Controller RespondWithJson Test Success")

}
