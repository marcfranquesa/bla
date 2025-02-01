package main

import (
	"fmt"
	"github.com/marcfranquesa/bla/pkg/db"
	"net/http"
)

func main() {
    conn, _ := db.Connect()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		results, _ := db.UrlById(conn, "43")
		fmt.Fprint(w, results)
	})

	http.ListenAndServe(":8080", nil)
}
