package llm

import (
	"fmt"
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
		err := fmt.Errorf("GROK provider not implemented yet")
		return nil, err
	default:
		err := fmt.Errorf("`%v` provider not supported", provider)
		return nil, err
	}
}
