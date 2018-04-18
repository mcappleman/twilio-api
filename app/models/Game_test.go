package models_test

import (
	"fmt"
	"testing"

	"github.com/mcappleman/twilio-api/app"
	"github.com/mcappleman/twilio-api/app/models"
)

var a app.App

func initTests() {

	a = app.App{}
	a.Init()

}

func TestGetGames(t *testing.T) {

	fmt.Println("Game Model GetGames Test Started")

	initTests()

	games, err := models.GetGames(a.BC.DB)

	if err != nil {
		t.Fail()
		return
	}

	fmt.Println("Games")
	fmt.Println(len(games))

	fmt.Println("Game Model GetGames Test Success")

}
