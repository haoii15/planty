package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var db *bun.DB // Global variable to hold the DB connection

const (
	host     = "localhost"
	port     = 61524
	user     = "haoii"
	password = ""
	dbname   = "haoii"

)



func init() {
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	fmt.Println(connectionString)
	// Open a connection to the database
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connectionString)))

	// Initialize bun with PostgreSQL dialect
	db = bun.NewDB(sqldb, pgdialect.New())

	// Verify the connection
	if err := db.Ping(); err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Successfully connected!")
}