package models

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/mgo.v2"
	"github.com/mcappleman/twilio-api/mongodb"
)

var db *mgo.Database

func initTests() {

	session, err := mongodb.NewSession(os.Getenv("DATABASE_URL"), os.Getenv("DATABASE_NAME"))
	if err != nil {
		panic(err)
	}

	db = session.Database()

}

func TestGetBucket(t *testing.T) {

	fmt.Println("Game Model GetBucket Test Started")

	initTests()

	games, err := GetBucket(db, 0, 101)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if len(games) == 0 {
		fmt.Println(len(games))
		t.Fail()
		return
	}

	fmt.Println("Game Model GetBucket Test Success")

}

func TestGetGames(t *testing.T) {

	fmt.Println("Game Model GetGames Test Started")

	games, err := GetGames(db)
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	if len(games) == 0 {
		fmt.Println(err)
		t.Fail()
		return
	}

	fmt.Println("Game Model GetGames Test Success")

}
