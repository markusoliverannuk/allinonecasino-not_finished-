package helper

import (
	"database/sql"
)

// Database pointer
var Db *sql.DB

// Initializes sqlite3 database - Opens if not created then also creates database at given path
func InitializeDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}
	return db
}

// Create tables for database if it does not exist, it creates a new one
func CreateDbTable(db *sql.DB) {
	//users table
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (id STRING PRIMARY KEY, name STRING, email STRING, password STRING, role INT, balance INT)`)
	if err != nil {
		panic(err)
	}

}
