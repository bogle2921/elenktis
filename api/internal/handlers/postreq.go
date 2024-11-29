package handlers

import (
	"log"
	"net/http"

	"github.com/bogle2921/elenktis/api/internal/models"
	"github.com/bogle2921/elenktis/api/internal/utils"
	"github.com/gin-gonic/gin"
)

func PostLoginPage(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse req."})
		log.Println(err)
		return
	}

	err = user.Validate()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		log.Println(err)
		return
	}
	token, err := utils.GenerateToken(user.Username, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logged in", "token": token})
}

func PostAddUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse req."})
		log.Println(err)
		return
	}
	//c.GetInt64("userId")

	err = user.AddUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Made user"})
}