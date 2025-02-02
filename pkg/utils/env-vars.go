package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	} else if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("Invalid port value %s: %v.", port, err)
	}
	return port
}

func GetDomain() string {
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		domain = fmt.Sprintf("localhost:%s", GetPort())
	}
	return domain
}
