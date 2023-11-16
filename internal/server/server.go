package server

import (
	"log"
	"net/http"

	"github.com/astrobounce/goqu/internal/storage"
)

type Config struct {
	ListenAddr string
	Store      storage.Storage
}

type Server struct {
	*Config
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{
		Config: cfg,
	}, nil
}

func (s *Server) Start() {
	if err := http.ListenAndServe(s.ListenAddr, s); err != nil {
		log.Fatalf("could not start server due to %s", err)
	} else {
		log.Printf("Starting Server on address %s", s.ListenAddr)
	}

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
}
