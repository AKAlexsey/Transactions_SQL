package main

import (
	"log"
	"net/http"

	"github.com/AKAlexsey/transactions-sql/controllers"
	"github.com/AKAlexsey/transactions-sql/models"
)

func main() {
	// Initialize the database
	db := models.InitDB()
	defer db.Close()

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
