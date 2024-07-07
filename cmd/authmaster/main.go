package main

import (
	"log"

	"github.com/BaneleJerry/AuthMaster/internal/config"
	"github.com/BaneleJerry/AuthMaster/internal/server"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	srv := server.NewServer(cfg)
	if err := srv.Start(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}
