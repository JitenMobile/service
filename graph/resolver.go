package graph

import (
	"context"
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
	PromptStore     *service.PromptStore
}

func NewDictionaryResolver(firestoreClient *firestore.Client, openaiClient *openai.Client) *Resolver {
	return &Resolver{
		DictionaryStore: service.NewDictionaryService(firestoreClient),
		LLMService:      service.NewLLMService(openaiClient),
		PromptStore:     service.NewPromptStore(),
	}
}

func (r *Resolver) ResolveWordQuery(ctx context.Context, word string) (*model.Word, error) {

	wordData, err := r.DictionaryStore.GetWord(ctx, word)
	if err == nil {
		return wordData, nil
	}
	if strings.HasSuffix(err.Error(), "not found") {
		prompt := r.PromptStore.GenerateDefinitionsPrompt()
		jsonDescription := r.PromptStore.JsonDescription()
		wordData, err := r.LLMService.StructuredOutput(ctx, word, prompt, jsonDescription)
		if err != nil {
			return nil, err
		}
		err = r.DictionaryStore.WriteWord(ctx, wordData)
		return wordData, err
	} else {
		return nil, err
	}
}
