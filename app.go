package main

import (
	"log"
	"net/http"

	"github.com/adelattia/fizzbuzz-rest-api/fizzbuzz"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/fizzbuzz", fizzbuzz.FizzBuzz).Methods("GET")
}
