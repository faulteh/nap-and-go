// Package main provides the web UI
// servers.go has the handlers for the server view
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"golang.org/x/oauth2"

	"github.com/faulteh/nap-and-go/config"
)

// Guild represents a Discord guild/server
type Guild struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Owner       bool   `json:"owner"`
	Permissions int64  `json:"permissions"`
}

// serversHandler handles the server view
func serversHandler(c *gin.Context) {
	// Get the user session
	session := sessions.Default(c)
	token := session.Get("token").(*oauth2.Token)

	// Retrieve list of servers for user from discord
	oauth2Config := config.LoadDiscordConfig().OAuth2Config()
	client := oauth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://discord.com/api/users/@me/guilds")
	if err != nil {
		log.Printf("Failed to get user info: %v\n", err)
		// For now just return an error
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user info",
			"msg":   err.Error(),
		})
		return
	}
	defer resp.Body.Close()		//nolint:errcheck

	if resp.StatusCode != http.StatusOK {
		log.Printf("Discord API responded with status %d\n", resp.StatusCode)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch servers"})
		return
	}

	var guilds []Guild
	if err := json.NewDecoder(resp.Body).Decode(&guilds); err != nil {
		log.Printf("Error decoding guilds response: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode servers"})
		return
	}

	// Filter guilds where the user has admin permissions
	var adminGuilds []Guild
	const ADMINISTRATOR = 0x00000008
	for _, guild := range guilds {
		if guild.Permissions&ADMINISTRATOR != 0 {
			adminGuilds = append(adminGuilds, guild)
		}
	}

	// Render the servers view
	c.HTML(http.StatusOK, "servers.html", gin.H{
		"Title": "Servers",
		"Servers": adminGuilds,
	})
}
