package controllers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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
