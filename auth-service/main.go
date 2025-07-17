package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("ShivaniRiddhisha") // Replace with a secure key

func GenerateJWT() (string, error) {
	// Define token claims
	claims := jwt.MapClaims{
		"username": "shivanigupta",
		"exp":      time.Now().Add(time.Hour * 4).Unix(), // Expires in 4 hour
		"iat":      time.Now().Unix(),                    // Issued at
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// JWT endpoint handler
func GenerateTokenHandler(c *gin.Context) {
	token, err := GenerateJWT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func main() {
	r := gin.Default()

	// Add the JWT generation endpoint
	r.GET("/generate-token", GenerateTokenHandler)

	// Start the server
	if err := r.Run(":8082"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
