// Package main provides the web UI
// servers.go has the handlers for the server view
package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	"github.com/faulteh/nap-and-go/discordapi"
)

// serversHandler handles the server view
func serversHandler(c *gin.Context) {
	// Get the user session
	session := sessions.Default(c)
	token := session.Get("token").(*oauth2.Token)

	// Retrieve list of servers for user and bot from discord
	userGuilds, err := discordapi.UserAdminGuildList(token)
	if err != nil {
		log.Printf("Failed to get user guilds: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user guilds",
			"msg":   err.Error(),
		})
		return
	}

	botGuilds, err := discordapi.BotGuildList()
	if err != nil {
		log.Printf("Failed to get bot guilds: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get bot guilds",
			"msg":   err.Error(),
		})
		return
	}

	// Mark the servers where the bot is present
	for i := range userGuilds {
		for _, botGuild := range botGuilds {
			if userGuilds[i].ID == botGuild.ID {
				userGuilds[i].HasBot = true
				break
			}
		}
	}

	// Render the servers view
	c.HTML(http.StatusOK, "servers.html", gin.H{
		"Title": "Servers",
		"UserServers": userGuilds,
	})
}
