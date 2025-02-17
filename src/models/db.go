package models

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitDB() *sql.DB {
	connStr := "postgres://postgres:postgres@db:5432/webapp?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func GetDB() *sql.DB {
	return db
}
