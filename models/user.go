// Package models represents the database models for the application.
// This model represents a Discord user associated with a specific server.
package models

import (
	"time"
)

// User represents a Discord user associated with a specific server.
type User struct {
	ID        uint      `gorm:"primaryKey"`                                // Primary key
	UserID    string    `gorm:"not null;index:user_server_idx,unique"`     // Discord User ID
	ServerID  string    `gorm:"not null;index:user_server_idx,unique"`     // Discord Server ID
	Username  string    `gorm:"not null"`                                  // Username
	Nickname  string    `gorm:""`                                          // Nickname
	Role      string    `gorm:""`                                          // Role
	JoinedAt  time.Time `gorm:"not null"`                                  // Join timestamp
	CreatedAt time.Time `gorm:"autoCreateTime"`                            // Created timestamp
	UpdatedAt time.Time `gorm:"autoUpdateTime"`                            // Updated timestamp
	DeletedAt *time.Time `gorm:"index"`                                    // Soft delete
	Server    Server `gorm:"foreignKey:ServerID;references:ID"`            // Server relation
}
