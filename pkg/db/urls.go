package db

import (
	"fmt"
	"log"
)

const table = "urls"

type URL struct {
	Id  string
	Url string
}

func GetAll() ([]URL, error) {
	query := fmt.Sprintf("SELECT id, url FROM %s", table)
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []URL
	for rows.Next() {
		var url URL
		if err := rows.Scan(&url.Id, &url.Url); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}
	return urls, rows.Err()
}

func InsertUrl(url URL) error {
	query := fmt.Sprintf("INSERT INTO %s (id, url) VALUES (?, ?)", table)
	_, err := conn.Exec(query, url.Id, url.Url)
	if err == nil {
		log.Printf("Inserted URL: %s with ID: %s into DB.", url.Url, url.Id)
	}
	return err
}

func UrlByID(id string) (URL, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", table)
	row := conn.QueryRow(query, id)

	var url URL
	err := row.Scan(&url.Id, &url.Url)
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
