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
	OwnerID     string    `gorm:"not null"`             // Discord user ID of the owner
	Region      string    `gorm:"not null"`             // Server region
	MemberCount int       `gorm:"not null"`             // Current member count
	CreatedAt   time.Time `gorm:"autoCreateTime"`       // Record creation time
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`       // Last update time
	DeletedAt   *time.Time `gorm:"index"`               // Soft delete timestamp

	Users []User `gorm:"foreignKey:ServerID;references:ID"` // Users in the server
}
