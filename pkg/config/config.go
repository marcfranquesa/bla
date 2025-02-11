package config

import (
	"os"
	"sync"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port   string
	Domain string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Load() (*Config, error) {
	var err error
	once.Do(func() {
		instance = &Config{}
		if err = instance.loadServer(); err != nil {
			return
		}
		if err = instance.loadDatabase(); err != nil {
			return
		}
	})
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (c *Config) loadServer() error {
	c.Server = ServerConfig{
		Port:   getEnv("SERVER_PORT", "8080"),
		Domain: getEnv("DOMAIN", "https://bla.cat"),
	}
	return nil
}

func (c *Config) loadDatabase() error {
	c.Database = DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "user"),
		Password: getEnv("DB_PASSWORD", "pwd"),
		Name:     getEnv("DB_NAME", "bla"),
	}
	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
