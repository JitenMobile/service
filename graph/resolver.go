package graph

import (
	"cloud.google.com/go/firestore"
	"github.com/jiten-mobile/service/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DictionaryStore *service.DictionaryStore
}

func NewFirestoreResolver(firestoreClient *firestore.Client) *Resolver {
	return &Resolver{
		DictionaryStore: service.NewDictionaryService(firestoreClient),
	}
}
