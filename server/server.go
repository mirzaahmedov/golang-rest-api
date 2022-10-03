package server

import (
	"api/models"
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

	if err := s.store.Open(); err != nil {
		return err
	}

	return http.ListenAndServe(s.config.PORT, s.router)
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

	s.router.PathPrefix("/register").HandlerFunc(func(w http.ResponseWriter, req *http.Request){
		user, err := s.store.RegisterUser(&models.User{
			Id: 100,
			Name: "From Golang",
		})
		if err != nil {
			log.Fatal(err)
		}

		bytes, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(201)
		w.Write(bytes)
	})
}