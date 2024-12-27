// Package web runs Gin based webserver for the application.
// auth.go has endpoints to handle authentication and session management
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	"github.com/faulteh/nap-and-go/config"
)

// AuthRequired is middleware to check if the user is authenticated
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			// User not logged in, redirect to login
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		// User is logged in; proceed to the next handler
		c.Next()
	}
}

// homeHandler redirects to login or servers page based on the session
func homeHandler(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		// Redirect to login page
		c.Redirect(http.StatusFound, "/login")
		return
	}
	// Redirect to servers page
	c.Redirect(http.StatusFound, "/servers/")
}

// loginPageHandler handles the home page rendering login page
func loginPageHandler(c *gin.Context) {
	// Render the home page
	c.HTML(http.StatusOK, "login.html", nil)
}

// loginRedirectHandler redirects the user to the Discord OAuth2 login page
func loginRedirectHandler(c *gin.Context) {
	oauth2Config := config.LoadDiscordConfig().OAuth2Config()
	url := oauth2Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

// discordCallbackHandler handles the Discord OAuth2 callback
func discordCallbackHandler(c *gin.Context) {
	state := c.Query("state")
	if state != "state-token" {
		// Redirect back to login page such that the login page
		c.Redirect(http.StatusFound, "/login")
		return
	}

	code := c.Query("code")
	if code == "" {
		// Redirect back to login page such that the login page
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// Exchange code for token
	oauth2Config := config.LoadDiscordConfig().OAuth2Config()
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Token exchange failed: %v\n", err)
		// Redirect back to login page such that the login page
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// Fetch user info
	client := oauth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://discord.com/api/users/@me")
	if err != nil {
		log.Printf("Failed to get user info: %v\n", err)
		// Redirect back to login page such that the login page
		c.Redirect(http.StatusFound, "/login")
		return
	}
	defer resp.Body.Close()		//nolint:errcheck

	var user map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Printf("Failed to parse user info: %v\n", err)
		// Redirect back to login page such that the login page
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// Store user data in session
	session := sessions.Default(c)
	session.Set("user", user)
	session.Set("token", token)
	if err := session.Save(); err != nil {
		log.Printf("Failed to save session: %v\n", err)
		// Redirect back to login page such that the login page
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// Redirect to the servers page
	c.Redirect(http.StatusFound, "/servers/")
}

// logoutHandler handles the logout
func logoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()	//nolint:errcheck
	// Redirect back to login page
	c.Redirect(http.StatusFound, "/login")
}