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

type Message struct {
	To		[]string
	From	[]string
	Body	string
}

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

	/*message := &Message{
		To: []string{"+14076342712"},
		From: []string{"+14077974748"},
		Body: "Unit Test",
	}*/

	form := url.Values{}
	form.Add("To", "[+14074362712]")
	form.Add("From", "[+14077974748]")
	form.Add("Body", "Unit Test")

	/*jsonMessage, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}*/

	r, _ := http.NewRequest("POST", "/games", strings.NewReader(form.Encode()))
	r.Form = form
	w := httptest.NewRecorder()

	a.BC.PostMessage(w, r)

	if w.Code < 200 || w.Code > 299 {
		t.Fail()
		return
	}

	fmt.Println(w.Body)

	fmt.Println("Game Controller PostMessage Test Success")

}
