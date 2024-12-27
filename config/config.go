// Package config provides the configuration for the application.
package config

import (
	"fmt"
	"os"

	"golang.org/x/oauth2"
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

// DiscordConfig holds the Discord configuration.
type DiscordConfig struct {
	clientID     string
	clientSecret string
	botToken     string
	redirectURL  string
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

// Define Discord's OAuth2 endpoint
var discordEndpoint = oauth2.Endpoint{
	AuthURL:  "https://discord.com/api/oauth2/authorize",
	TokenURL: "https://discord.com/api/oauth2/token",
}

// LoadDiscordConfig loads the Discord configuration from environment variables.
func LoadDiscordConfig() *DiscordConfig {
	return &DiscordConfig{
		clientID:     getEnv("DISCORD_CLIENT_ID", ""),
		clientSecret: getEnv("DISCORD_CLIENT_SECRET", ""),
		botToken:     getEnv("DISCORD_BOT_TOKEN", ""),
		redirectURL:  getEnv("DISCORD_REDIRECT_URL", ""),
	}
}

// OAuth2Config Extend the DiscordConfig struct with a method OAuth2Config that returns an oauth2.Config.
func (c *DiscordConfig) OAuth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
		Endpoint:     discordEndpoint,
		RedirectURL:  c.redirectURL,
		Scopes:       []string{"identify", "email", "guilds"},
	}
}

// LoadSessionStoreSecret Load the session store secret from an environment variable.
func LoadSessionStoreSecret() string {
	return getEnv("SESSION_STORE_SECRET", "")
}
