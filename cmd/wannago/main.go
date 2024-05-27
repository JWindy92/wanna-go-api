package main

import (
	"log"

	"github.com/JWindy92/wanna-go-api/internal/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("cant initialize zap logger: %v", err)
	}
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
	r := gin.Default()

	// handlers.Pong()
	// Setup routes
	r.GET("/ping", handlers.Pong)

	// // Start the server
	r.Run(":8080")
}
