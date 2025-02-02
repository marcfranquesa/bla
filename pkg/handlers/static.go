package handlers

import (
	"log"
	"net/http"
	"path/filepath"
)

func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	webDir, err := filepath.Abs("./web")
	if err != nil {
		log.Printf("Failed to get absolute path to web directory: %v.", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(webDir)).ServeHTTP(w, r)
}
