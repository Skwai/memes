package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
)

type App struct {
	Router 	*mux.Router
}

func (a *App) Initialize() {
  a.Router = mux.NewRouter()
  a.initializeRoutes()
}

func (a *App) Run(addr string) {
  http.ListenAndServe(addr, a.Router)
}

func (a *App) initializeRoutes() {
  a.Router.HandleFunc("/memes", a.getMemes).Methods("GET")
}

func (a *App) getMemes(w http.ResponseWriter, r *http.Request) {
  m, err := getRandomMemes(2)

  if err != nil {
    respondWithError(w, http.StatusInternalServerError, "Oh noes")
  }

  respondWithJSON(w, http.StatusOK, m)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
  respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}
