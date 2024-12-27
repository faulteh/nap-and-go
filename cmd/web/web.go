// Package web runs Gin based webserver for the application.
package main

import (
	"log"
	
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the web server
	router := gin.Default()
	
	// Load the templates
	router.LoadHTMLGlob("templates/*")

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
	// Define the routes
	router.GET("/", homeHandler)
	router.GET("/test", testHandler)
}
