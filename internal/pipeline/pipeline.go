// Package pipeline
package pipeline

import (
	"strings"

	"github.com/mabd-dev/doc-gen/internal/ollama"
)

type Pipeline struct {
	Ollama *ollama.Client
}

func (p Pipeline) Analyze(code, prompt string) (string, error) {
	finalPrompt := strings.Replace(prompt, "{{FUNCTION}}", code, 1)

	return p.Ollama.Generate(finalPrompt)
}

func (p Pipeline) GenerateDoc(
	analysis, functionSignature, prompt string,
) (string, error) {
	finalPrompt := strings.Replace(prompt, "{{ANALYSIS}}", analysis, 1)
	finalPrompt = strings.Replace(finalPrompt, "{{SIGNATURE}}", functionSignature, 1)

	return p.Ollama.Generate(finalPrompt)
}
