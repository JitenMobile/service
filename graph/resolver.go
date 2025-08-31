package graph

import (
	"context"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/jiten-mobile/service/graph/model"
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

func (r *Resolver) ResolveWordQuery(ctx context.Context, word string) (*model.Word, error) {

	wordData, err := r.DictionaryStore.GetWord(ctx, word)
	// Should allow "not found" error pass
	if err == nil {
		return wordData, nil
	}

	// Only generate new word content when the error is "not found"
	if strings.HasPrefix(err.Error(), service.NotFoundPrefix) {
		wordData, err := r.LLMService.StructuredWord(ctx, word)
		if err != nil {
			log.Printf("Error generating word '%s' from LLMService: '%v", word, err)
			return nil, err
		}
		err = r.DictionaryStore.WriteWord(ctx, wordData)
		// don't return error to the query, as the data is ready for user
		if err != nil {
			log.Printf("Error writing word data to database: %v", err)
		}
		return wordData, nil
	}
	return nil, err
}

func (r *Resolver) ResolveTranslationQuery(ctx context.Context, word string, targetLang string) (*model.Translation, error) {

	translationData, err := r.DictionaryStore.GetTranslation(ctx, word, targetLang)
	if err == nil {
		return translationData, nil
	}

	wordData, err := r.DictionaryStore.GetWord(ctx, word)
	if err != nil {
		log.Printf("Error reading word data from databse '%v'", err)
		return nil, err
	}

	definitions := wordData.Definitions
	newTranslationData, err := r.LLMService.StructuredTranslation(ctx, targetLang, definitions)
	if err != nil {
		log.Printf("Error generating definition for word '%s' from LLMService: '%v", word, err)
		return nil, err
	}

	err = r.DictionaryStore.WriteTranslation(ctx, word, targetLang, newTranslationData)
	if err != nil {
		// don't return error to the query, as the data is ready for user
		log.Printf("Error writing translations to database: %v", err)
	}
	return newTranslationData, nil
}
