package server

import (
	"api/handlers"
	"api/store"
	"net/http"
	"os"

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

	s.setup()

	if err := s.store.Open(); err != nil {
		return err
	}

	if len(os.Getenv("PORT")) > 0 {
		s.config.PORT = "0.0.0.0:" + os.Getenv("PORT")
	}

	return http.ListenAndServe(s.config.PORT, s.router)
}

func (s *Server)use(path string, handler handlers.RouteHandler) {
	handler(path, s.router)
}
func (s *Server) setup(){
	
	s.use("/user", handlers.UserHandler)

}