package core

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

func InitFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	projectId := os.Getenv("PROJECT_ID")
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}
	return client, nil
}
