package app_test

import (
	"fmt"
	"testing"

	"github.com/mcappleman/twilio-api/app"
)

var a app.App

func TestInit(t *testing.T) {

	fmt.Println("App Test Init Started")
	a = app.App{}
	err := a.Init()
	if err != nil {
		t.Fail()
	}
	fmt.Println("App Test Init Success")

}

func TestInitRoutes(t *testing.T) {

	fmt.Println("App Test Init Routes Started")
	a = app.App{}
	err := a.Init()
	if err != nil {
		t.Fail()
	}
	a.InitRoutes()
	fmt.Println("App Test Init Routes Success")

}
