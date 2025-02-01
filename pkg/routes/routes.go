package routes

import (
	"github.com/marcfranquesa/bla/pkg/handlers"
	"net/http"
)

func SetupRoutes() {
	http.Handle("/", handlers.ServeStaticFiles())

	http.HandleFunc("/l/", handlers.Redirect)
}
