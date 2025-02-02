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

	id, err := generateId()
	if err != nil {
		return failureResponse
	}

	err = db.AddUrl(db.URL{Id: id, Url: url})
	if err != nil {
		log.Printf("Failed to add URL '%s': %v.", url, err)
		return failureResponse
	}
	log.Printf("Added URL '%s' with ID '%s'.", url, id)

	domain := utils.GetDomain()
	return Response{
		Message: fmt.Sprintf("%s/l/%s", domain, id),
		Status:  "success",
	}
}

func generateId() (string, error) {
	var id string
	var exists bool
	attempts := 10
	for i := 0; i < attempts; i++ {
		id = utils.GenerateId()
		exists, _ = db.IDExists(id)
		if !exists {
			return id, nil
		}
	}
	return "", fmt.Errorf("failed to generate ID after %d attempts", attempts)
}
