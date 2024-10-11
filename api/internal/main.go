package main

import (
	"github.com/bogle2921/elenktis/api/internal/database"
	"github.com/bogle2921/elenktis/api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	server := gin.Default()

	server.GET("/", handlers.GetLoginPage)
	server.POST("/", handlers.PostLoginPage)

	server.Run(":8080")
}
