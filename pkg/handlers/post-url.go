package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/marcfranquesa/bla/pkg/config"
	"github.com/marcfranquesa/bla/pkg/db"
	"github.com/marcfranquesa/bla/pkg/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	Url string `json:"url"`
}

func PostUrl(w http.ResponseWriter, r *http.Request, cfg config.ServerConfig) {
	var url string
	contentType := r.Header.Get("Content-Type")

	if contentType == "application/json" {
		var request Request
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&request); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		url = request.Url
	} else {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body.", http.StatusBadRequest)
			return
		}
		url = strings.TrimSpace(string(body))
	}

	if !validateURL(url) {
		http.Error(w, "Invalid URL format.", http.StatusBadRequest)
		return
	}

	response, err := getResponse(url, cfg)
	if err != nil {
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, response)
}

func validateURL(urlStr string) bool {
	_, err := url.ParseRequestURI(urlStr)
	return err == nil
}

func getResponse(url string, cfg config.ServerConfig) (string, error) {
	id := utils.GenerateId(url, 4)

	used, err := db.IsIDUsed(id, url)
	if err != nil {
		log.Printf("Error checking if ID '%s' is used: %v", id, err)
		return "Failed to check ID usage", err
	}

	if used {
		log.Printf("ID '%s' already exists with a different URL", id)
		return "Error: ID already in use, try a longer length", nil
	}

	inserted, err := db.IsIDInserted(id)
	if err != nil {
		log.Printf("Error checking if ID '%s' is inserted: %v", id, err)
		return "Failed to check insertion status", err
	}

	if !inserted {
		err = db.InsertUrl(db.URL{Id: id, Url: url})
		if err != nil {
			log.Printf("Error inserting URL '%s': %v", url, err)
			return "Failed to insert URL", err
		}
	}

	return fmt.Sprintf("%s/l/%s", cfg.Domain, id), nil
}
