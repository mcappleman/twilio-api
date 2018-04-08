package app_test

import (
	"os"
	"testing"

	"github.com/mcappleman/twilio-api/app"
)

var a app.App

func TestMain(m *testing.M) {

	a = app.App{}
	a.Init()

	os.Exit(0)

}
