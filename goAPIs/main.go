package main

import (
	"log"
	"net/http"

	"example.com/goAPI/models"
	"example.com/goAPI/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	server.GET("/users", getUsers)

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getUsers(c *gin.Context) {
	users, err := utils.LoadFromFile[models.User]("data/Users.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load users"})
		return
	}

	var userList []models.UserResponse
	for _, user := range users {
		userList = append(userList, models.UserResponse{
			Name:   user.Name,
			Email:  user.Email,
			Phone:  user.Phone,
			Status: user.Status,
		})
	}
	c.JSON(http.StatusOK, gin.H{"users": userList})
}
