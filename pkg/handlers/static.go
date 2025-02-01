package handlers

import (
	"log"
	"net/http"
	"path/filepath"
)

func ServeStaticFiles() http.Handler {
	webDir, err := filepath.Abs("../../web")
	if err != nil {
		log.Fatalf("Failed to get absolute path to web directory: %v", err)
	}

    log.Printf("Serving static files from: %s", webDir)

	return http.FileServer(http.Dir(webDir))
}

