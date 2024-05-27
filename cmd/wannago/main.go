package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Setup routes
	r.GET("/ping", handlers.Ping)

	// Start the server
	r.Run(":8080")
}
