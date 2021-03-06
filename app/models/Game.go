package models

import (
	// "fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Game struct {
	Id                  bson.ObjectId `json:"id" bson:"_id"`
	Date                time.Time     `json:"date" bson:"date"`
	HomeTeam            bson.ObjectId `json:"home_team" bson:"home_team"`
	AwayTeam            bson.ObjectId `json:"away_team" bson:"away_team"`
	HomeRuns            int           `json:"home_runs" bson:"home_runs"`
	AwayRuns            int           `json:"away_runs" bson:"away_runs"`
	Status              string        `json:"status" bson:"status"`
	NumberFireFavorite bson.ObjectId `json:"number_fire_favorite" bson:"number_fire_favorite"`
	NumberFireOdds      float64       `json:"number_fire_odds" bson:"number_fire_odds"`
}

const collectionName = "games"

func GetBucket(db *mgo.Database, min float64, max float64) ([]Game, error) {

	gameList := []Game{}

	err := db.C(collectionName).Find(
			bson.M{
				"status": "Final",
				"number_fire_odds":
					bson.M{
						"$gte": min,
						"$lt": max,
					},
				},
			).All(&gameList)
	if err != nil {
		return nil, err
	}

	return gameList, nil

}

func GetGames(db *mgo.Database) ([]Game, error) {

	return GetBucket(db, 0, 101)

}
