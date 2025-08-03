package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/goAPI/configs"
	"example.com/goAPI/internal/handlers"
	"example.com/goAPI/internal/middleware"
	"example.com/goAPI/internal/repositories"
	"example.com/goAPI/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config := configs.LoadConfig()

	// Debug logging
	log.Printf("Server configuration: Host=%s, Port=%s", config.Server.Host, config.Server.Port)
	log.Printf("Data file path: %s", config.Database.FilePath)

	// Initialize dependencies
	userRepo := repositories.NewUserRepository(config.Database.FilePath)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Setup Gin router
	router := gin.New()

	// Add middleware
	router.Use(middleware.Logger())
	router.Use(middleware.ErrorHandler())
	router.Use(gin.Recovery())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"service": "User API",
			"version": "1.0.0",
		})
	})

	// API routes
	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUserByID)
		}
	}

	// Start server
	serverAddr := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	log.Printf("Starting server on %s", serverAddr)

	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
