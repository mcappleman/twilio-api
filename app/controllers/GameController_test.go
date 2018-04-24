package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestGetGames(t *testing.T) {

	fmt.Println("Game Controller GetGames Test Started")

	r, _ := http.NewRequest("GET", "/games", nil)
	w := httptest.NewRecorder()

	bc.GetGames(w, r)

	if w.Code != 200 {
		fmt.Println("W Code: " + strconv.Itoa(w.Code))
		t.Fail()
		return
	}

	var body ListResponse
	decoder := json.NewDecoder(w.Body)
	err := decoder.Decode(&body)
	if err != nil {
		fmt.Println("Decoder Error")
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

func TestGetBuckets(t *testing.T) {

	fmt.Println("Game Controller getBuckets Test Started")

	returnMessage, err := getBuckets(bc.DB)
	if err != nil {
		t.Fail()
		return
	}

	if returnMessage == "" {
		t.Fail()
		return
	}

	fmt.Println(returnMessage)

	fmt.Println("Game Controller getBuckets Test Success")

}

func TestFillBucket(t *testing.T) {

	fmt.Println("Game Controller fillBucket Test Started")

	var err error
	bucket := buckets{}
	bucket.bucket50_55, err = fillBucket(bc.DB, 50, 55)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	fmt.Println("Game Controller fillBucket Test Success")

}
