package service

import "github.com/openai/openai-go/v2"

type LLMService struct {
	client *openai.Client
}

func NewLLMService(client *openai.Client) *LLMService {
	return &LLMService{client: client}
}
