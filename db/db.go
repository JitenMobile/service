package db

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

func InitClient() (*firestore.Client, error) {
	ctx := context.Background()
	projectId := os.Getenv("PROJECT_ID")
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}
	defer client.Close()
	return client, nil
}
