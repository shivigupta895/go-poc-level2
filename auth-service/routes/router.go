package routes

import (
	"auth-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/token", handlers.TokenHandler)

	return r
}
