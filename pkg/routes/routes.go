package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/marcfranquesa/bla/pkg/config"
	"github.com/marcfranquesa/bla/pkg/handlers"
)

func SetupRoutes(cfg config.ServerConfig) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Forwarded-For")
		if ip != "" {
			ip = strings.Split(ip, ",")[0]
		} else {
			ip = r.RemoteAddr
		}
		log.Printf("Client IP: %s, Method: %s, URL: %s, Proto: %s\n", ip, r.Method, r.URL, r.Proto)

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
