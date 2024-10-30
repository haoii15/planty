package api

import (
	"planty/db"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)


func CreateUser(c *gin.Context) {
	// Create a new user
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Call the createUser function from the db package
	err := db.CreateUser(firstName, lastName, username, password)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "User created successfully!",
		})
	}
}