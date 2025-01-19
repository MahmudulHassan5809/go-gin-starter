package main

import (
	"gin_starter/src/core/db"
	"gin_starter/src/core/middlewares"
	"gin_starter/src/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Initialize the database
	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Set up the Gin router
	router := gin.Default()

	
	router.Use(middlewares.ResponseMiddleware())
	router.Use(middlewares.RequestResponseLogger())

	// Register all routes
	routes.RegisterRoutes(router, db)
	

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	log.Printf("Starting server on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
