package routes

import (
	"net/http"

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

	http.HandleFunc("/l/", handlers.Redirect)
}
