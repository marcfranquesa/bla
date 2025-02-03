package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/marcfranquesa/bla/pkg/db"
	"github.com/marcfranquesa/bla/pkg/utils"
	"log"
	"net/http"
)

type Request struct {
	Url string `json:"url"`
}

func PostUrl(w http.ResponseWriter, r *http.Request) {
	var request Request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	response := getResponse(request.Url)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, response)
}

func getResponse(url string) string {
	id := utils.GenerateId(url, 3)

	used, err := db.IsIDUsed(id, url)
	if err != nil {
		log.Printf("Failed to add URL '%s': %v.", url, err)
		return "Error"
	}
	if used {
		log.Printf("ID '%s' exists with a different URL than '%s'", id, url)
		return "Error: try with a longer length"
	}

	inserted, err := db.IsIDInserted(id)
	if !inserted {
		db.AddUrl(db.URL{Id: id, Url: url})
	}

	domain := utils.GetDomain()
	return fmt.Sprintf("%s/l/%s", domain, id)
}
