package main

import (
	"log"
	"net/http"
	"time"

	"github.com/marcfranquesa/bla/pkg/config"
	"github.com/marcfranquesa/bla/pkg/db"
	"github.com/marcfranquesa/bla/pkg/routes"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v.", err)
	}

	for err := db.Connect(cfg.Database); err != nil; err = db.Connect(cfg.Database) {
		log.Printf("Failed to connect to database. Retrying... Error: %v.", err)
		time.Sleep(2 * time.Second)
	}
	log.Printf("Successfully connected to database.")
	defer db.Close()

	routes.SetupRoutes(cfg.Server)

	port := cfg.Server.Port
	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %v.", err)
	}
}
