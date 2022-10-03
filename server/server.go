package server

import (
	"api/store"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	config *Config
	router *mux.Router
	store *store.Store
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		router: mux.NewRouter(),
		store: store.New(),
	}
}

func (s *Server)Start() error {

	s.bindHandlers()

	if err := http.ListenAndServe(s.config.PORT, s.router); err != nil {
		return err
	}
	return nil
}

func (s *Server) bindHandlers(){
	type Person struct {
		Name string `json:"name"`
	}

	me := Person{
		Name: "Bekzod",
	}

	result, err := json.Marshal(&me)
	if err != nil {
		log.Fatal(err)
	}

	s.router.PathPrefix("/json").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(201)
		w.Write(result)
	}).Methods("GET")

	s.router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request){
		w.Write([]byte("Hello"))
	}).Methods("GET")
}