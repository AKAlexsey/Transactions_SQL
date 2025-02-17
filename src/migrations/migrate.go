package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Database connection string
	connStr := "postgres://postgres:postgres@127.0.0.1:5432/transactions-sql?sslmode=disable"

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Read migration file
	migrationSQL, err := os.ReadFile("src/migrations/000_create_database.sql")
	migrationSQL, err := os.ReadFile("src/migrations/001_create_users_table.sql")
	if err != nil {
		log.Fatal("Error reading migration file:", err)
	}

	// Execute migration
	_, err = db.Exec(string(migrationSQL))
	if err != nil {
		log.Fatal("Error executing migration:", err)
	}

	fmt.Println("Migration completed successfully!")
}
