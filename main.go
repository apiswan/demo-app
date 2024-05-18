package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Chat struct {
	Message string `json:"message" binding:"required"`
	Room    string `json:"room" binding:"required"`
}

func main() {
	// Create a new Gin router
	r := gin.Default()

	var chat = map[string]interface{}{
		"data": map[string]interface{}{
			"username":     "luffy",
			"message":      "what's up guys!",
			"id":           53443,
			"read-receipt": true,
		},
	}

	// Handle GET request
	r.GET("/chat-services/user/chat", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	
		// Check if the request method is OPTIONS (preflight request)
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusOK)
			return
		}
		c.JSON(http.StatusOK, chat)
	})

	// Handle POST request
	r.POST("/chat-services/user/chat", func(c *gin.Context) {
		// Set CORS headers to allow requests from all origins
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	
		// Check if the request method is OPTIONS (preflight request)
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusOK)
			return
		}
		// Now handle the POST request as before
		var chat Chat
		if err := c.BindJSON(&chat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		c.JSON(http.StatusOK, map[string]string{"id": uuid.New().String(), "message": chat.Message, "room": chat.Room})
	})
	

	// Run the server on port 8080
	r.Run(":8080")
}
