// Package pipeline
package pipeline

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mabd-dev/doc-gen/internal/ollama"
)

type Pipeline struct {
	Ollama *ollama.Client
}

func (p Pipeline) Analyze(code, prompt string) (string, error) {
	finalPrompt := strings.Replace(prompt, "{{FUNCTION}}", code, 1)

	return p.Ollama.GenerateWithModel(finalPrompt, p.Ollama.BaseModel)
}

func (p Pipeline) GenerateDoc(
	analysis, signature, prompt string,
) (string, error) {
	finalPrompt := strings.Replace(prompt, "{{ANALYSIS}}", analysis, 1)
	finalPrompt = strings.Replace(finalPrompt, "{{SIGNATURE}}", signature, 1)

	return p.Ollama.GenerateWithModel(finalPrompt, p.Ollama.BaseModel)
}

// GetDocsOnly filter output from [GenerateDoc] function and get only the documentation part
func (p Pipeline) GetDocsOnly(docs string) (string, error) {
	kdocRegex := regexp.MustCompile(`/\*\*(.|[\r\n])*?\*/`)
	matches := kdocRegex.FindAllString(docs, -1)

	if len(matches) > 0 {
		return matches[0], nil
	}

	err := fmt.Errorf("could not find kdoc in llm response")
	return "", err
}

func (p Pipeline) PolishDoc(doc, prompt string) (string, error) {
	finalPrompt := strings.Replace(prompt, "{{KDOC}}", doc, 1)
	return p.Ollama.GenerateWithModel(finalPrompt, p.Ollama.DocPolishModel)
}
