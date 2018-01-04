package main

import (
	"github.com/gorilla/mux"
)

type App struct {
	Router 	*mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {}