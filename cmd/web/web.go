// Package web runs Gin based webserver for the application.
package main

import (
	"log"
	"encoding/gob"

	"github.com/faulteh/nap-and-go/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func main() {
	// Initialize the web server
	router := gin.Default()
	
	// Register the routes
	registerRoutes(router)
	
	// Run the web server
	addr := ":8080"
	log.Printf("Starting web server on %s", addr)
	err := router.Run(addr)
	if err != nil {
		log.Fatalf("failed to start web server: %v", err)
	}
}

// registerRoutes registers the routes for the web server.
func registerRoutes(router *gin.Engine) {
	// Register map[string]interface{} with gob as we use it for user session data
	gob.Register(map[string]interface{}{})
	// Register oauth2.Token with gob as we use it for user session data
	gob.Register(&oauth2.Token{})

	// Set up session store to store login sessions
	store := cookie.NewStore([]byte(config.LoadSessionStoreSecret()))
	router.Use(sessions.Sessions("nap-and-go", store))
	
	// Load the templates
	router.LoadHTMLGlob("templates/*.html")

	// Static files
	router.Static("/static", "./static")

	// Define the routes
	router.GET("/", homeHandler)

	// Auth routes
	router.GET("/login", loginPageHandler)
	router.GET("/auth/login", loginRedirectHandler)
	router.GET("/auth/callback", discordCallbackHandler)
	router.GET("/logout", logoutHandler)

	authenticated := router.Group("/servers")
	authenticated.Use(AuthRequired())

	authenticated.GET("/", serversHandler)
	authenticated.GET("/:id", dashboardHandler)
}
