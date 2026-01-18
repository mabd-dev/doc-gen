package llm

import (
	"fmt"
	"os"
	"strings"
)

type Client interface {
	GetBaseModel() string
	GetDocPolishModel() string

	Generate(prompt string) (string, error)
	GenerateWithModel(prompt, model string) (string, error)
}

func NewClient(provider string) (Client, error) {
	switch strings.ToLower(strings.TrimSpace(provider)) {
	case "ollama":
		return OllamaClient{
			BaseURL:        "http://localhost:11434",
			BaseModel:      "qwen-kdoc",
			DocPolishModel: "llama-kdoc:latest",
		}, nil
	case "grok":
		return OpenAIClient{
			BaseURL:   "https://api.groq.com/openai/v1",
			APIKey:    os.Getenv("GROQ_API_KEY"),
			BaseModel: "qwen/qwen3-32b",
		}, nil
	default:
		err := fmt.Errorf("`%v` provider not supported", provider)
		return nil, err
	}
}
