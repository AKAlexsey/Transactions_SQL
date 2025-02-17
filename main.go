package main

import (
	"log"
	"net/http"

	"github.com/AKAlexsey/transactions-sql/src/controllers"
	"github.com/AKAlexsey/transactions-sql/src/models"
)

func main() {
	// Initialize the database
	db := models.InitDB()
	defer db.Close()
	// Check if the connection is successful

	err := db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}
	log.Println("Main. Successfully connected to PostgreSQL!")

	// Initialize controllers
	homeController := controllers.NewHomeController()

	// Routes
	http.HandleFunc("/", homeController.Index)

	// Start server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
