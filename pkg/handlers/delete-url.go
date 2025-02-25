package handlers

import (
	"fmt"
	"net/http"

	"github.com/marcfranquesa/bla/pkg/db"
)

func DeleteURL(w http.ResponseWriter, r *http.Request, id string) {
	db_token, err := db.TokenByID(id)
	if err != nil {
		http.Error(w, "Failed to retrieve token", http.StatusInternalServerError)
		return
	}
	if db_token != r.Header.Get("Authorization") {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	err = db.DeleteURL(id)
	if err != nil {
		http.Error(w, "Failed to delete URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "URL deleted")
}
