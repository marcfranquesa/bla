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

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func PostUrl(w http.ResponseWriter, r *http.Request) {
	var request Request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getResponse(request.Url))
}

func getResponse(url string) Response {
	failureResponse := Response{
		Message: "Unable to shorten URL",
		Status:  "failure",
	}

	id := utils.GenerateId(url, 3)

	used, err := db.IsIDUsed(id, url)
	if err != nil {
		log.Printf("Failed to add URL '%s': %v.", url, err)
		return failureResponse
	}
	if used {
		log.Printf("ID '%s' exists with a different URL than '%s'", id, url)
		return failureResponse
	}

	inserted, err := db.IsIDInserted(id)
	if !inserted {
		db.AddUrl(db.URL{Id: id, Url: url})
	}

	domain := utils.GetDomain()
	return Response{
		Message: fmt.Sprintf("%s/l/%s", domain, id),
		Status:  "success",
	}
}
