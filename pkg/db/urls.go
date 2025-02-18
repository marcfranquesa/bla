package db

import (
	"fmt"
	"log"
)

const table = "urls"

type URL struct {
	Id       string
	Url      string
	Verified uint8
}

func GetAllURLs() ([]URL, error) {
	query := fmt.Sprintf("SELECT id, url, verified FROM %s", table)
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []URL
	for rows.Next() {
		var url URL
		if err := rows.Scan(&url.Id, &url.Url, &url.Verified); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}
	return urls, rows.Err()
}

func InsertUrl(id string, urlStr string) error {
	query := fmt.Sprintf("INSERT INTO %s (id, url) VALUES (?, ?)", table)
	_, err := conn.Exec(query, id, urlStr)
	if err == nil {
		log.Printf("Inserted URL: %s with ID: %s into DB.", id, urlStr)
	}
	return err
}

func DeleteURL(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id =?", table)
	_, err := conn.Exec(query, id)
	return err
}

func VerifyURL(id string) error {
	query := fmt.Sprintf("UPDATE %s SET verified = 1 WHERE id =?", table)
	_, err := conn.Exec(query, id)
	return err
}

func UrlByID(id string) (URL, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", table)
	row := conn.QueryRow(query, id)

	var url URL
	err := row.Scan(&url.Id, &url.Url, &url.Verified)
	return url, err
}

func IsIDInserted(id string) (bool, error) {
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE id = ?", table)
	rows, err := conn.Query(query, id)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}

func IsIDUsed(id string, url string) (bool, error) {
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE id = ? AND url <> ?", table)
	rows, err := conn.Query(query, id, url)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}
