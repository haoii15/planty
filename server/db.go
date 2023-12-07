package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)



func createPlantsTable() {
	_, err := db.Exec(`
		CREATE TABLE plants (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255),
			description TEXT,
			ip VARCHAR(255)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'plants' created!")
}

func createDataTable() {
	_, err := db.Exec(`
		CREATE TABLE data (
			id SERIAL PRIMARY KEY,
			value INT,
			time TIMESTAMP,
			plant INT REFERENCES plants(id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'data' created!")
}

func tableExists(tableName string) bool {
	rows, err := db.Query("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = $1)", tableName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var exists bool
	for rows.Next() {
		if err := rows.Scan(&exists); err != nil {
			log.Fatal(err)
		}
	}

	return exists
}

func insertPlant(name string, description string, ip string) {
	_, err := db.Exec("INSERT INTO plants (name, description, ip) VALUES ($1, $2, $3)", name, description, ip)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Plant record inserted!")
}

func insertData(value int, plantID int) {
	_, err := db.Exec("INSERT INTO data (value, time, plant) VALUES ($1, $2, $3)", value, time.Now(), plantID)
	if err != nil {
		fmt.Println("Error in insertData")
		log.Fatal(err)
	}

	fmt.Println("Data record inserted!")
}

func getPlantIDByIP(ip string) (int, error) {
	var plantID int
	err := db.QueryRow("SELECT id FROM plants WHERE ip = $1", ip).Scan(&plantID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case when no plant is found with the given IP
			return 0, err
		} else {
			log.Fatal(err)
			return 0, err
		}
	}

	fmt.Println("Plant ID: ", plantID)

	return plantID, nil
}

func getNewestDataForEachPlant() []map[string]interface{} {
	query := `
		SELECT p.id AS plant, p.name AS plant_name, d.id AS data_id, d.value, d.time
		FROM plants p
		JOIN data d ON p.id = d.plant
		WHERE d.time = (
			SELECT MAX(time)
			FROM data
			WHERE plant = p.id
		)
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// List to store the result
	resultList := []map[string]interface{}{}

	fmt.Println("Newest data for each plant:")
	for rows.Next() {
		var plantID int
		var plantName string
		var dataID int
		var value int
		var time time.Time

		err := rows.Scan(&plantID, &plantName, &dataID, &value, &time)
		if err != nil {
			log.Fatal(err)
		}

		// Create a dictionary for the current plant
		plantData := map[string]interface{}{
			"name":  plantName,
			"value": value,
			"time": time,
		}

		// Append the dictionary to the list
		resultList = append(resultList, plantData)

		fmt.Printf("Plant Name: %s, Value: %d\n", plantName, value)
	}

	return resultList
}

