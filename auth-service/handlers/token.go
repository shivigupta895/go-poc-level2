package handlers

import (
	"auth-service/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type TokenRequest struct {
	Username string `json:"username" binding:"required"`
}

func TokenHandler(c *gin.Context) {
	jwtUserName, err := utils.GetSecret("JWT_USER_NAME")
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Failed to get JWT_USER_NAME"})
		return
	}

	claims := jwt.MapClaims{
		"username": jwtUserName,
		"exp":      time.Now().Add(time.Hour * 4).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecretKey, err := utils.GetSecret("JWT_SECRET_KEY")
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Failed to get JWT_SECRET_KEY"})
		return
	}

	tokenString, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
