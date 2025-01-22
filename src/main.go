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
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		sqlDB, _ := database.DB()
		sqlDB.Close()
	}()

	router := gin.Default()
	router.Use(middlewares.ErrorHandlingMiddleware())
	router.Use(middlewares.RequestResponseLogger())
	routes.RegisterRoutes(router, database)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
