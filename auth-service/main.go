package main

import (
	"auth-service/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8082") // Starts server on localhost:8082
}
