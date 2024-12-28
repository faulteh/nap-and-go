// Package db provides a database connection and helper functions to interact with the database.
// server.go has functions to interact with the server table in the database.
package db

import (
	"errors"
	"log"

	"gorm.io/gorm"

	"github.com/faulteh/nap-and-go/discordtypes"
	"github.com/faulteh/nap-and-go/models"
)

// SyncGuilds synchronizes the servers in the database with the servers in the Discord API
func SyncGuilds(guilds []discordtypes.Guild) error {
	// Sync the guilds with the database
	conn := GetDB()
	for _, guild := range guilds {
		var server models.Server
		if err := conn.Where("id = ?", guild.ID).First(&server).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Create a new server
				server = models.Server{
					ID:             guild.ID,
					Name:           guild.Name,
					Icon:           guild.Icon,
					Owner:          guild.Owner,
					Permissions:    guild.Permissions,
					PermissionsNew: guild.PermissionsNew,
				}
				if err := conn.Create(&server).Error; err != nil {
					log.Printf("Failed to create server: %v\n", err)
					return err
				}
			} else {
				log.Printf("Failed to query server: %v\n", err)
				return err
			}
		} else {
			// Update the server
			server.Name = guild.Name
			server.Icon = guild.Icon
			server.Owner = guild.Owner
			server.Permissions = guild.Permissions
			server.PermissionsNew = guild.PermissionsNew
			if err := conn.Save(&server).Error; err != nil {
				log.Printf("Failed to update server: %v\n", err)
				return err
			}
		}
	}
	return nil
}
