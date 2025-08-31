package service

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/jiten-mobile/service/graph/model"
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/shared"
)

type LLMService struct {
	client    *openai.Client
	prompts   *PromptStore
	chatModel shared.ChatModel
}

func NewLLMService(client *openai.Client) *LLMService {
	return &LLMService{
		client:    client,
		prompts:   NewPromptStore(),
		chatModel: openai.ChatModelGPT4_1Mini,
	}
}

// Docs
func (llm *LLMService) StructuredWord(ctx context.Context, word string) (*model.Word, error) {

	jsonSchema := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        "definition_generation",
		Description: openai.String(llm.prompts.generationJsonDescription),
		Schema:      JsonTypeOf(reflect.TypeOf(model.Word{})),
	}

	resp, err := llm.client.Chat.Completions.New(
		ctx,
		openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(llm.prompts.generateDefinitionsPrompt),
				openai.UserMessage(word),
			},
			ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
				OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{JSONSchema: jsonSchema},
			},
			Model: llm.chatModel,
		},
	)

	if err != nil {
		return nil, err
	}

	var data model.Word
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// Docs
func (llm *LLMService) StructuredTranslation(ctx context.Context, targetLang string, definitions []*model.Definition) (*model.Translation, error) {

	wordDataJson, err := json.Marshal(definitions)
	if err != nil {
		return nil, err
	}

	jsonSchema := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        "definition_translation",
		Description: openai.String(llm.prompts.translationJsonDescription),
		Schema:      JsonTypeOf(reflect.TypeOf(model.Translation{})),
	}

	resp, err := llm.client.Chat.Completions.New(
		ctx,
		openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(llm.prompts.GetTranslationPrompt(targetLang)),
				openai.UserMessage(string(wordDataJson)),
			},
			ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
				OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{JSONSchema: jsonSchema},
			},
			Model: llm.chatModel,
		},
	)

	if err != nil {
		return nil, err
	}

	var data model.Translation
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &data); err != nil {
		return nil, err
	}
	return &data, nil
}
