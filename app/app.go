package app

import (
	"fmt"
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

	// log.Println("App Init Method Begin")
	// fmt.Println("App Init Method Begin")
	db, err := mongodb.NewSession(os.Getenv("DATABASE_URL"), os.Getenv("DATABASE_NAME"))
	if err != nil { return err }
	// log.Println("App Init Mongo Connected")
	// fmt.Println("App Init Mongo Connected")
	a.Router = mux.NewRouter()
	a.BC = controllers.Init(db.Database())
	// log.Println("App Init Method End")
	// fmt.Println("App Init Method End")

	return nil

}

func (a *App) InitRoutes() {

	fmt.Println("Init Routes")
	a.Router.HandleFunc("/", a.BC.Index).Methods("GET")
	a.Router.HandleFunc("/games", a.BC.GetGames).Methods("GET")
	a.Router.HandleFunc("/message", a.BC.PostMessage).Methods("POST")

	a.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil { return err }
		fmt.Println(t)
		return nil
	})

}

func (a *App) Run(addr string) {

	fmt.Println("Running on port :8080")
    log.Fatal(http.ListenAndServe(addr, a.Router))

}
