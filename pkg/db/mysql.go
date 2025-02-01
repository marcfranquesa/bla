package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"os"
	"sync"
)

var (
	conn *sql.DB
	once sync.Once
)

type URL struct {
	Id  string
	Url string
}

func Connect() error {
	once.Do(func() {
		cfg := mysql.Config{
			User:   os.Getenv("MYSQL_USER"),
			Passwd: os.Getenv("MYSQL_PASSWORD"),
			Net:    "tcp",
			Addr:   os.Getenv("MYSQL_ADDRESS"),
			DBName: os.Getenv("MYSQL_DATABASE"),
		}

		conn, _ = sql.Open("mysql", cfg.FormatDSN())
	})

	return conn.Ping()
}

func AddUrl(url URL) error {
	_, err := conn.Exec("INSERT INTO urls (id, url) VALUES (?, ?)", url.Id, url.Url)
	if err != nil {
		return err
	}
	return nil
}

func UrlById(id string) (URL, error) {
	var url URL
	row := conn.QueryRow("SELECT * FROM urls WHERE id = ?", id)
	err := row.Scan(&url.Id, &url.Url)
	return url, err
}
