// Package main provides the web UI
// dashboard.go has the handlers for the dashboard view and configuration of the bot
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// dashboardHandler handles the dashboard view
func dashboardHandler(c *gin.Context) {
	// Render the dashboard view
	c.HTML(http.StatusOK, "dashboard.html", nil)
}
