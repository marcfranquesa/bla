package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"os"
)

type URL struct {
	id  string
	url string
}

func Connect() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("MYSQL_ADDRESS"),
		DBName: os.Getenv("MYSQL_DATABASE"),
	}

	var conn *sql.DB
	var err error
	conn, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func AddUrl(conn *sql.DB, url URL) error {
	_, err := conn.Exec("INSERT INTO urls (id, url) VALUES (?, ?)", url.id, url.url)
	if err != nil {
		return err
	}
	return nil
}

func UrlById(conn *sql.DB, id string) (URL, error) {
	var url URL
	row := conn.QueryRow("SELECT * FROM urls WHERE id = ?", id)
	err := row.Scan(&url.id, &url.url)
	return url, err
}
