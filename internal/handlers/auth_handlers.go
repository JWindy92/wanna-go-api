package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Pong(c *gin.Context) {
	zap.L().Info("Pong")
	c.String(http.StatusOK, "Pong")
}

func Login(c *gin.Context) {
	zap.L().Info("Login")
	c.String(http.StatusOK, "Login")
}

func Register(c *gin.Context) {
	zap.L().Info("Register")
	c.String(http.StatusOK, "Register")
}
