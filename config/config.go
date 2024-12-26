// Package config provides the configuration for the application.
package config

import (
	"fmt"
	"os"
)

// DBConfig holds the database configuration.
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// LoadDBConfig loads the database configuration from environment variables.
func LoadDBConfig() *DBConfig {
	return &DBConfig{
		Host:     getEnv("DB_HOST", "postgres"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "napandgo"),
		Password: getEnv("DB_PASSWORD", "napandgo"),
		DBName:   getEnv("DB_NAME", "napandgo"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

// getEnv gets an environment variable with a default fallback.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// DSN returns the PostgreSQL Data Source Name.
func (c *DBConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}
