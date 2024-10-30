package main

import (
	"planty/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.New()
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000","http://localhost:3001"} // Add the origin of your React app
	router.Use(cors.New(config))

	// Define the routes
	router.POST("/user", api.CreateUser)

	router.Run(":61942")
}