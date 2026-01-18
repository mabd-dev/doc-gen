package llm

import (
	"fmt"
	"os"
	"strings"

	clitypes "github.com/mabd-dev/doc-gen-ai/internal/cli_types"
)

type Client interface {
	ShouldDoDocsPolishig() bool

	GetBaseModel() string
	GetDocPolishModel() string

	Generate(prompt string) (string, error)
	GenerateWithModel(prompt, model string) (string, error)
}

func NewClient(
	provider string,
	baseURLFlag *string,
	baseModelFlag *string,
	polishDocsFlag clitypes.OptionalBool,
) (Client, error) {
	switch strings.ToLower(strings.TrimSpace(provider)) {
	case "ollama":
		url := DefaultOllamaBaseURL
		if baseURLFlag != nil && len(*baseURLFlag) > 0 {
			url = *baseURLFlag
		}

		baseModel := DefaultOllamaBaseModel
		if baseModelFlag != nil && len(*baseModelFlag) > 0 {
			baseModel = *baseModelFlag
		}

		polishDocs := DefaultOllamaPolishDocs
		if polishDocsFlag.IsSet {
			polishDocs = polishDocsFlag.Value
		}

		return OllamaClient{
			BaseURL:        url,
			BaseModel:      baseModel,
			DocPolishModel: "llama-kdoc:latest",
			PolishDocs:     polishDocs,
		}, nil
	case "groq":
		url := DefaultOpenAIBaseURL
		if len(*baseURLFlag) > 0 {
			url = *baseURLFlag
		}

		baseModel := DefaultOpenAIBaseModel
		if len(*baseModelFlag) > 0 {
			baseModel = *baseModelFlag
		}

		polishDocs := DefaultOpenAIPolishDocs
		if polishDocsFlag.IsSet {
			polishDocs = polishDocsFlag.Value
		}

		return OpenAIClient{
			BaseURL:    url,
			APIKey:     os.Getenv("GROQ_API_KEY"),
			BaseModel:  baseModel,
			PolishDocs: polishDocs,
		}, nil
	default:
		err := fmt.Errorf("`%v` provider not supported", provider)
		return nil, err
	}
}
