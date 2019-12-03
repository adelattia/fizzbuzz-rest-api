package main

import (
	"github.com/adelattia/fizzbuzz-rest-api/core"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	Port string
}

func (a *App) Initialize() {
	a.Router = a.initializeRoutes()
	a.Port = "8080"
	log.Printf("listening on port %s..", a.Port)
	log.Fatal(http.ListenAndServe(":"+a.Port, a.Router))
}

func (a *App) initializeRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz", core.FizzBuzz).Methods("GET")
	return r
}
