package routes

import (
	"net/http"
	"strings"

	"github.com/marcfranquesa/bla/pkg/config"
	"github.com/marcfranquesa/bla/pkg/handlers"
)

func SetupRoutes(cfg config.ServerConfig) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.ServeStaticFiles(w, r)
		} else if r.Method == http.MethodPost {
			handlers.PostUrl(w, r, cfg)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/l/", func(w http.ResponseWriter, r *http.Request) {
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

		if r.Method == http.MethodGet {
			handlers.Redirect(w, r, id)
		} else if r.Method == http.MethodDelete {
			handlers.DeleteURL(w, r, id)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
}
