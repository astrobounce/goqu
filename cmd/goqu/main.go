package main

import (
	"log"

	"github.com/astrobounce/goqu/internal/server"
	"github.com/astrobounce/goqu/internal/storage"
)

func main() {
	cfg := &server.Config{
		ListenAddr: ":3000",
		Store:      storage.NewMemoryStore(),
	}

	server, err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("could not start server due to %s", err)
	}
	server.Start()
}
