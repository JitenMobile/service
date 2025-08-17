package service

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/jiten-mobile/service/graph/model"
	"github.com/openai/openai-go/v2"
)

type LLMService struct {
	client *openai.Client
}

func NewLLMService(client *openai.Client) *LLMService {
	return &LLMService{client: client}
}

func (llm *LLMService) StructuredOutput(ctx context.Context, word string, systemPrompt string, jsonDescription string) (*model.Word, error) {

	jsonSchema := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        "definition_generation",
		Description: openai.String(jsonDescription),
		Schema:      JsonTypeOf(reflect.TypeOf(model.Word{})),
	}

	resp, err := llm.client.Chat.Completions.New(
		ctx,
		openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(systemPrompt),
				openai.UserMessage(word),
			},
			ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
				OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{JSONSchema: jsonSchema},
			},
			Model: openai.ChatModelGPT4_1Mini,
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
