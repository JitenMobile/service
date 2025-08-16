package core

import (
	"os"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func InitOpenaiClient() *openai.Client {
	openaiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(option.WithAPIKey(openaiKey))
	return &client
}
