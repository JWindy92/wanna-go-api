package main

import (
	"github.com/JWindy92/wanna-go-api/internal/handlers"
	"github.com/JWindy92/wanna-go-api/internal/logging"
	"github.com/JWindy92/wanna-go-api/internal/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	logger = logging.GetLogger()
	defer logger.Sync()
	logger.Info("Starting application")

	db, err := utils.Connect()
	if err != nil {
		logger.Sugar().Errorf("Error opening database connection: %v\n", err)
	}
	defer db.Close()

	// zap.ReplaceGlobals(logger)
	r := gin.Default()

	r.GET("/ping", handlers.Pong)
	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)

	r.Run(":8080")
}
