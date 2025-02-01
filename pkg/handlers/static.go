package handlers

import (
	"log"
	"net/http"
	"path/filepath"
)

func ServeStaticFiles() http.Handler {
	webDir, err := filepath.Abs("./web")
	if err != nil {
		log.Printf("Failed to get absolute path to web directory: %v", err)
		return nil
	}

	return http.FileServer(http.Dir(webDir))
}
