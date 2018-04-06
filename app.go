package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mcappleman/twilio-api/mongodb"
)

type App struct {
	Router	*mux.Router
	DB		*mongodb.DatabaseSession
}

func (a *App) Init() error {

	r.DB, err = mongodb.NewSession(os.Getenv("MONGO_URL"), os.Getenv("DATABASE_NAME"))
	if err != nil { return err }
	a.Router = mux.NewRouter()

}

func (a *App) Run(addr string) {

    log.Fatal(http.ListenAndServe(":8080", a.Router))

}
