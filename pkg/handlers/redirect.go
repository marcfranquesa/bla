package handlers

import (
	"github.com/marcfranquesa/bla/pkg/db"
	"net/http"
	"strings"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	parts := strings.Split(path, "/")
	if len(parts) > 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id := parts[2]
	if id == "" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	url, _ := db.UrlById(id)

	http.Redirect(w, r, url.Url, http.StatusFound)
}
