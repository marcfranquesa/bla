package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/marcfranquesa/bla/pkg/config"
)

var (
	conn *sql.DB
	once sync.Once
)

func Connect(cfg config.DatabaseConfig) error {
	var err error
	once.Do(func() {
		sql_cfg := mysql.Config{
			User:   cfg.User,
			Passwd: cfg.Password,
			Net:    "tcp",
			Addr:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			DBName: cfg.Name,
		}

		conn, err = sql.Open("mysql", sql_cfg.FormatDSN())
	})

	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	return conn.Ping()
}

func Close() error {
	if conn != nil {
		return conn.Close()
	}
	return nil
}
