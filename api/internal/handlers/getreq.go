package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLoginPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Got login page"})
}