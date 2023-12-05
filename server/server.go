package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var units = make(map[string]string)
var db *sql.DB

const (
	host     = "localhost"
	port     = 61524
	user     = "haoii"
	password = ""
	dbname   = "haoii"

)

func get_all(c *gin.Context) {
	// var dataList []map[string]string

	// for key, value := range units {
	// 		entry := make(map[string]string)
	// 		entry["name"] = key
	// 		entry["value"] = value
	// 		dataList = append(dataList, entry)
	// }

	datalist := getNewestDataForEachPlant()

	c.JSON(http.StatusOK, datalist)
}

func store_data(c *gin.Context) {
	// id := c.Query("id")
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error reading request body"})
		return
	}

	ip := c.ClientIP()
	
	id, err := getPlantIDByIP(ip)
	if err != nil{
		if err == sql.ErrNoRows {
			fmt.Println("No plant found with IP: " + ip, "Inserting new plant")
			insertPlant("plant "+ strconv.Itoa(rand.Intn(100) + 1),  "default description", ip)
		} else {
			fmt.Println("Error in query")
			log.Fatal(err)
		}
	}

	val, err := strconv.Atoi(string(body))
	if err != nil {
		fmt.Println("Error in strconv.Atoi")
		log.Fatal(err)
	}
	insertData(val, id)

	// units[ip] = string(body)
	
	// if _, exists := units[ip]; exists {
	// 	c.String(http.StatusOK, "thanks")
	// } else {
	// 	c.String(http.StatusOK, "welcome")
	// }
}


func init() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open a connection to the database
	_db, err := sql.Open("postgres", connectionString)
	if err != nil {
			log.Fatal(err)
	}
	
	err = _db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db = _db
	
	// Check if tables exist
	if !tableExists("plants"){
		createPlantsTable()
	}
	
	if !tableExists("data") {
		createDataTable()
	}
	
	
	fmt.Println("Successfully connected!")
}


func main() {
	defer db.Close()
	router := gin.New()
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Add the origin of your React app
	router.Use(cors.New(config))


	router.GET("/", get_all)			// curl localhost:8080/message
	router.POST("/", store_data)			// curl -X PUT localhost:8080/message -d "msg=tull" -H "Content-Type: application/x-www-form-urlencoded"

	router.Run(":61942")
}