// Package models represents the database models for the application.
// This model represents a Discord server.
package models

import (
	"time"
)

// Server represents a Discord server.
type Server struct {
	ID          string    `gorm:"primaryKey"`            // Discord server ID (Snowflake)
	Name        string    `gorm:"not null"`             // Server name

	Icon        string    `gorm:"type:text"`    // Server icon URL
	Banner	    string    `gorm:"type:text"`    // Server banner URL

	Owner       bool      `gorm:"not null;default:false"` // Is the bot owner of the server
	Permissions int64     `gorm:"not null"`             // Bot permissions in the server
	PermissionsNew string `gorm:"not null"`              // New permissions in the server
	Features    []string  `gorm:"type:text[]"`          // Server features

	CreatedAt   time.Time `gorm:"autoCreateTime"`       // Record creation time
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`       // Last update time
	DeletedAt   *time.Time `gorm:"index"`               // Soft delete timestamp

	Users []User `gorm:"foreignKey:ServerID;references:ID"` // Users in the server
}
