package handlers

import (
	"net/http"

	"github.com/JWindy92/wanna-go-api/internal/logging"
	"github.com/JWindy92/wanna-go-api/internal/models"
	"github.com/JWindy92/wanna-go-api/internal/utils"
	"github.com/gin-gonic/gin"
)

var logger = logging.GetLogger()

// type LoginRequest struct {
// 	Username string `json:"username"`
// }
// type RegistrationRequest struct {
// 	Username string `json:"username"`
// }

func Pong(c *gin.Context) {
	logger.Info("Pong")
	c.String(http.StatusOK, "Pong")
}

func Login(c *gin.Context) {
	logger.Info("Login")

	var msg models.LoginRequest

	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "received",
		"message": msg,
	})
}

func Register(c *gin.Context) {
	logger.Info("Register")
	var msg models.RegistrationRequest

	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.CreateUser(msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "received",
		"message": msg,
	})
}
