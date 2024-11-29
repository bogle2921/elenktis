package main

import (
	"github.com/bogle2921/elenktis/api/internal/database"
	"github.com/bogle2921/elenktis/api/internal/handlers"
	"github.com/bogle2921/elenktis/api/internal/middleware"

	"github.com/gin-gonic/gin"
)

/*
Protected routes:
GET /home
GET /adduser
POST /adduser
GET /systems
POST /systems
*/
func main() {
	database.InitDB()
	server := gin.Default()

	// unauthenicated endpoints
	server.GET("/", handlers.GetLoginPage)
	server.POST("/", handlers.PostLoginPage)

	// protected endpoints
	authenicated := server.Group("/")
	authenicated.Use(middleware.Authenticate)
	//authenicated.POST("/systems", handlers.PostAddSystem)

	server.Run(":8080")
}
