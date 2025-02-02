package routes

import (
	"github.com/marcfranquesa/bla/pkg/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.ServeStaticFiles(w, r)
		} else if r.Method == http.MethodPost {
			handlers.PostUrl(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/l/", handlers.Redirect)
}
