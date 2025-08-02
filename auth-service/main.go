package main

import (
	"auth-service/routes"
	// "auth-service/utils"
)

func main() {
	// utils.LoadEnvVariables()

	r := routes.SetupRouter()
	r.Run(":8082") // Starts server on localhost:8082
}
