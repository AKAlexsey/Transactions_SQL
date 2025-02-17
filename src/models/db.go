package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() *sql.DB {
	connStr := "postgres://postgres:postgres@db:5432/transactions-sql?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to PostgreSQL!")

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func GetDB() *sql.DB {
	return db
}
