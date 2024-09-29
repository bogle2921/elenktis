package handlers

import (
	"log"
	"net/http"

	"github.com/bogle2921/elenktis/api/internal/models"
	"github.com/gin-gonic/gin"
)

func GetLoginPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logged in"})
}

func PostLoginPage(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse req."})
		log.Println(err)
		return
	}

	user.ID = 1
	c.JSON(http.StatusCreated, gin.H{"message": "Logged in"})
}
