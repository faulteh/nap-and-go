// Package web runs Gin based webserver for the application.
// auth.go has endpoints to handle authentication
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// homeHandler handles the home page rendering index.html
func homeHandler(c *gin.Context) {
	// Render the home page
	c.HTML(http.StatusOK, "index.html", nil)
}

func testHandler(c *gin.Context) {
	// Render the home page
	c.String(http.StatusOK, "Hello World")
}
