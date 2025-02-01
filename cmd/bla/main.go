package main

import (
	"github.com/marcfranquesa/bla/pkg/db"
	"github.com/marcfranquesa/bla/pkg/routes"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := getPort()

	conn, _ := db.Connect()
	routes.SetupRoutes(conn)

	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	} else if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("Invalid port value %s: %v", port, err)
	}
	return port
}
