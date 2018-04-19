package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGetGames(t *testing.T) {

	fmt.Println("Game Controller GetGames Test Started")

	r, _ := http.NewRequest("GET", "/games", nil)
	w := httptest.NewRecorder()

	bc.GetGames(w, r)

	if w.Code != 200 {
		t.Fail()
		return
	}

	var body ListResponse
	decoder := json.NewDecoder(w.Body)
	err := decoder.Decode(&body)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	fmt.Println("Game Controller GetGames Test Success")

}

func TestPostMessage(t *testing.T) {

	fmt.Println("Game Controller PostMessage Test Started")

	form := url.Values{}
	form.Add("To", "[+14074362712]")
	form.Add("From", "[+14077974748]")
	form.Add("Body", "Unit Test")

	r, _ := http.NewRequest("POST", "/games", strings.NewReader(form.Encode()))
	r.Form = form
	w := httptest.NewRecorder()

	bc.PostMessage(w, r)

	if w.Code < 200 || w.Code > 299 {
		t.Fail()
		return
	}

	fmt.Println("Game Controller PostMessage Test Success")

}

func TestSendMessage(t *testing.T) {

	fmt.Println("Game Controller sendMessage Test Started")

	message := "Unit Test SendMessage"
	to := "+14077974748"
	from := "+4074362712"

	err := sendMessage(message, to, from)
	if err != nil {
		t.Fail()
		return
	}

	fmt.Println("Game Controller sendMessage Test Success")

}

func TestGetBucket(t *testing.T) {

	fmt.Println("Game Controller getBucket Test Started")

	list, err := getBucket(50, 55)
	if err != nil {
		t.Fail()
		return
	}

	if len(list) == 0 {
		t.Fail()
		return
	}

	fmt.Println("Game Controller getBucket Test Success")

}
