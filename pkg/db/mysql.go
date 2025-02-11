package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/marcfranquesa/bla/pkg/config"
	"log"
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

func Connect(cfg config.DatabaseConfig) error {
	once.Do(func() {
		sql_cfg := mysql.Config{
			User:   cfg.User,
			Passwd: cfg.Password,
			Net:    "tcp",
			Addr:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			DBName: cfg.Name,
		}

		conn, _ = sql.Open("mysql", sql_cfg.FormatDSN())
	})

	return conn.Ping()
}

func Close() error {
    if conn != nil {
        return conn.Close()
    }
    return nil
}

func InsertUrl(url URL) error {
	_, err := conn.Exec("INSERT INTO urls (id, url) VALUES (?, ?)", url.Id, url.Url)
	if err == nil {
		log.Printf("Inserted URL: %s with ID: %s into DB.", url.Url, url.Id)
	}
	return err
}

func UrlByID(id string) (URL, error) {
	var url URL
	row := conn.QueryRow("SELECT * FROM urls WHERE id = ?", id)
	err := row.Scan(&url.Id, &url.Url)
	return url, err
}

func IsIDInserted(id string) (bool, error) {
	rows, err := conn.Query("SELECT 1 FROM urls WHERE id = ?", id)
	defer rows.Close()

	if err != nil {
		return false, err
	}

	if rows.Next() {
		return true, nil
	}

	return false, nil
}

func IsIDUsed(id string, url string) (bool, error) {
	rows, err := conn.Query("SELECT 1 FROM urls WHERE id = ? AND url <> ?", id, url)
	defer rows.Close()

	if err != nil {
		return false, err
	}

	if rows.Next() {
		return true, nil
	}

	return false, nil
}
