package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if os.Getenv("ENV") == "" || os.Getenv("ENV") == "dev" {
		err := godotenv.Load()

		if err != nil {
			log.Println(`{"message":"Couldn't load .env file. Proceeding with system environment.", "service":"order", "severity":"INFO"}`)
		} else {
			log.Println(`{"message":"Environment file loaded.", "service":"order", "severity":"INFO"}`)
		}
	}
}

func GetSecret(secretName string) (string, error) {
	ctx := context.Background()

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Printf(`{"message":"Failed to create secret manager client: %v", "service":"order", "severity":"ERROR"}`, err)
		return "", fmt.Errorf("failed to create secret manager client: %v", err)
	}
	defer client.Close()
	log.Println(`{"message":"Client created, Accessing Request", "service":"order", "severity":"INFO"}`)

	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", os.Getenv("PROJECT_NUMBER"), secretName),
	}

	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Printf(`{"message":"failed to access secret: %v", "service":"order", "severity":"ERROR"}`, err)
		return "", fmt.Errorf("failed to access secret: %v", err)
	}
	log.Printf(`{"message":"%s Successfully Retrieved", "service":"order", "severity":"INFO"}`, secretName)

	return string(result.Payload.Data), nil
}
