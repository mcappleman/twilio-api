package controllers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mcappleman/twilio-api/app"
)

var a app.App

func initTest() {

	a = app.App{}
	a.Init()

}

func TestIndex(t *testing.T) {

	fmt.Println("Base Controller Index Test Started")
	initTest()

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	a.BC.Index(w, r)

	if (w.Code != 200) {
		t.Fail()
	}

	var body map[string]string
	json.Unmarshal(w.Body.Bytes(), &body)

	if (body["message"] != "Hello Twilio World!") {
		t.Fail()
	}

	fmt.Println("Base Controller Index Test Success")

}

func TestRespondWithJson(t *testing.T) {

	fmt.Println("Base Controller RespondWithJson Test Started")

	w := httptest.NewRecorder()

	a.BC.RespondWithJson(w, 200, map[string]string{"message": "Test"})

	if (w.Code != 200) {
		t.Fail()
	}

	var body map[string]string
	json.Unmarshal(w.Body.Bytes(), &body)

	if (body["message"] != "Test") {
		t.Fail()
	}

	fmt.Println("Base Controller RespondWithJson Test Success")

}
