package main

import (
	"github.com/marcfranquesa/bla/pkg/db"
	"github.com/marcfranquesa/bla/pkg/routes"
	"github.com/marcfranquesa/bla/pkg/utils"
	"log"
	"net/http"
	"time"
)

func main() {
	port := utils.GetPort()

	for err := db.Connect(); err != nil; err = db.Connect() {
		log.Printf("Failed to connect to database. Retrying... Error: %v.", err)
		time.Sleep(2 * time.Second)
	}
	log.Printf("Successfully connected to database.")

	routes.SetupRoutes()

	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %v.", err)
	}
}
