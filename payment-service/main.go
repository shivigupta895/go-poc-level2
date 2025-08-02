package main

import (
	"log"
	"payment-service/routes"
	"payment-service/utils"
)

func main() {
	utils.LoadEnvVariables()
	r := routes.SetupRouter()
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
