package main

import (
	"backend_go/api/repositories"
	"backend_go/api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	err := repositories.InitDB("words.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Set up the Gin router
	r := gin.Default()
	routes.SetupRoutes(r)

	// Run the server
	r.Run(":8080")
}
