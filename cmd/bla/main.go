package main

import (
	"github.com/marcfranquesa/bla/pkg/db"
	"github.com/marcfranquesa/bla/pkg/routes"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	port := getPort()

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

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	} else if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("Invalid port value %s: %v.", port, err)
	}
	return port
}
