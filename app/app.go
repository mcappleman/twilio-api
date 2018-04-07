package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mcappleman/twilio-api/mongodb"
	"github.com/mcappleman/twilio-api/app/controllers"
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

func (a *App) InitRoutes() {

	a.Router.HandleFunc("/games", a.GetGames).Methods("GET")

}

func (a *App) Run(addr string) {

    log.Fatal(http.ListenAndServe(":8080", a.Router))

}

func RespondWithError(w http.ResponseWriter, code int, message string) {

	RespondWithJson(w, code, map[string]string{"error": message})

}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {

	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}
