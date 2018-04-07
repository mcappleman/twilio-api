package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mcappleman/twilio-api/mongodb"
	"github.com/mcappleman/twilio-api/app/controllers"
)

type App struct {
	Router	*mux.Router
	BC		controllers.BaseController
}

func (a *App) Init() error {

		db, err := mongodb.NewSession(os.Getenv("MONGO_URL"), os.Getenv("DATABASE_NAME"))
	if err != nil { return err }
	a.Router = mux.NewRouter()
	a.BC = controllers.Init(db.Database())

	return nil

}

func (a *App) InitRoutes() {

	a.Router.HandleFunc("/games", a.BC.GetGames).Methods("GET")

}

func (a *App) Run(addr string) {

    log.Fatal(http.ListenAndServe(":8080", a.Router))

}
