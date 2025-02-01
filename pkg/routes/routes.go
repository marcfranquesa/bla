package routes

import (
	"database/sql"
	"fmt"
	"github.com/marcfranquesa/bla/pkg/db"
	"github.com/marcfranquesa/bla/pkg/handlers"
	"net/http"
)

func SetupRoutes(conn *sql.DB) {
	http.Handle("/", handlers.ServeStaticFiles())

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		results, _ := db.UrlById(conn, "43")
		fmt.Fprint(w, results)
	})
}
