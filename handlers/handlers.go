package handlers

import (
	"api/store"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteHandler func (string, *mux.Router)

func UserHandler(path string, router *mux.Router) {
	route := router.PathPrefix(path).Subrouter()
	store := store.New()
	
	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		users, err := store.GetAllUsers()
		if err != nil {
			log.Fatal(err)
		}

		bytes, err := json.Marshal(users)
		if err != nil {
			log.Fatal()
		}

		w.Write(bytes)
	}).Methods("GET")

	route.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello From Golang"))
	}).Methods("POST")
}