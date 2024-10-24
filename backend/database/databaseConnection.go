package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver import
)

var DB *sql.DB

// CreateDBConnection initializes and returns a MySQL connection, creating the DB if it doesn't exist
func CreateDBConnection() (*sql.DB, error) {
	// Connection string with parseTime=true
	dsn := "root:@tcp(127.0.0.1:3306)/go_db"

	// Open the connection to MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Failed to connect to MySQL:", err)
		return nil, err
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		log.Println("Failed to ping MySQL database:", err)
		return nil, err
	}

	return db, nil
}