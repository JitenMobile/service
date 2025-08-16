package graph

import (
	"cloud.google.com/go/firestore"
	"github.com/jiten-mobile/service/service"
	"github.com/openai/openai-go/v2"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DictionaryStore *service.DictionaryStore
	LLMService      *service.LLMService
}

func NewDictionaryResolver(firestoreClient *firestore.Client, openaiClient *openai.Client) *Resolver {
	return &Resolver{
		DictionaryStore: service.NewDictionaryService(firestoreClient),
		LLMService:      service.NewLLMService(openaiClient),
	}
}
