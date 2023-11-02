package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var units = make(map[string]string)

func get_all(c *gin.Context) {
	var dataList []map[string]string

	for key, value := range units {
			entry := make(map[string]string)
			entry["name"] = key
			entry["value"] = value
			dataList = append(dataList, entry)
	}
	for key, value := range units {
			entry := make(map[string]string)
			entry["name"] = key+key
			entry["value"] = value
			dataList = append(dataList, entry)
	}
	for key, value := range units {
			entry := make(map[string]string)
			entry["name"] = key+key+key
			entry["value"] = value
			dataList = append(dataList, entry)
	}

	c.JSON(http.StatusOK, dataList)
}

func store_data(c *gin.Context) {
	name := c.Query("name")
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error reading request body"})
		return
	}
	fmt.Println(string(body))
	fmt.Println(name)
	units[name] = string(body)
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}


func main() {
	router := gin.New()
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Add the origin of your React app
	router.Use(cors.New(config))


	router.GET("/", get_all)			// curl localhost:8080/message
	router.POST("/", store_data)			// curl -X PUT localhost:8080/message -d "msg=tull" -H "Content-Type: application/x-www-form-urlencoded"

	router.Run(":61942")
}