package controllers_test

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

	a.BC.GetGames(w, r)

	if (w.Code != 200) {
		t.Fail()
		return
	}

	var body map[string]string
	json.Unmarshal(w.Body.Bytes(), &body)

	fmt.Println("Body")
	fmt.Println(body)
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

	a.BC.PostMessage(w, r)

	if w.Code < 200 || w.Code > 299 {
		t.Fail()
		return
	}

	fmt.Println("Game Controller PostMessage Test Success")

}
