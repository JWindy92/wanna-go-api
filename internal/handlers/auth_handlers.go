package handlers

import (
	"net/http"

	"github.com/JWindy92/wanna-go-api/internal/logging"
	"github.com/gin-gonic/gin"
)

var logger = logging.GetLogger()

func Pong(c *gin.Context) {
	logger.Info("Pong")
	c.String(http.StatusOK, "Pong")
}

func Login(c *gin.Context) {
	logger.Info("Login")
	c.String(http.StatusOK, "Login")
}

func Register(c *gin.Context) {
	logger.Info("Register")
	c.String(http.StatusOK, "Register")
}
