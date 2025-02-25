package handlers

import (
	"net/http"

	"github.com/marcfranquesa/bla/pkg/db"
)

func Redirect(w http.ResponseWriter, r *http.Request, id string) {
	url, _ := db.UrlByID(id)
	http.Redirect(w, r, url, http.StatusFound)
}
