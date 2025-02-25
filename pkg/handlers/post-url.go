package handlers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/marcfranquesa/bla/pkg/config"
	"github.com/marcfranquesa/bla/pkg/db"
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

	response, token, err := getResponse(url, cfg)
	if err != nil {
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	if token != "" {
		w.Header().Set("X-Token", token)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, response)
}

func validateURL(urlStr string) bool {
	_, err := url.ParseRequestURI(urlStr)
	return err == nil
}

func getResponse(url string, cfg config.ServerConfig) (string, string, error) {
	id := generateId(url, 4)

	used, err := db.IsIDUsed(id, url)
	if err != nil {
		log.Printf("Error checking if ID '%s' is used: %v", id, err)
		return "Failed to check ID usage", "", err
	}

	if used {
		log.Printf("ID '%s' already exists with a different URL", id)
		return "Error: ID already in use, try a longer length", "", nil
	}

	inserted, err := db.IsIDInserted(id)
	if err != nil {
		log.Printf("Error checking if ID '%s' is inserted: %v", id, err)
		return "Failed to check insertion status", "", err
	}

	var token string
	if !inserted {
		token, err = tokenURLSafe(32)
		if err != nil {
			log.Printf("Error generating token: %v", err)
			return "Failed to generate token", "", err
		}

		err = db.InsertUrl(id, url, token)
		if err != nil {
			log.Printf("Error inserting URL '%s': %v", url, err)
			return "Failed to insert URL", "", err
		}
	}

	return fmt.Sprintf("%s/l/%s", cfg.Domain, id), token, nil
}

func generateId(s string, length int) string {
	data := []byte(s)

	hash := sha256.New()
	hash.Write(data)
	hashedData := hash.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(hashedData)[:length]
}

func tokenURLSafe(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
