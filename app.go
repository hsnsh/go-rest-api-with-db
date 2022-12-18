package main

import (
	"database/sql"
	c "go-rest-api-with-db/controllers"
	. "go-rest-api-with-db/helpers"
	// tom: for Initialize
	"log"
	"net/http"
	// tom: go get required
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	Logger *log.Logger
}

// Public functions

func (a *App) Initialize() {

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

// Private Functions

func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/", a.handleIndex).Methods("GET")

	//Register routes of controllers
	c.NewProductController().InitializeRoutes(a.Router)
}

func (a *App) handleIndex(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, http.StatusOK, "Welcome Product API")
}
