// Package pipeline
package pipeline

import (
	"strings"

	"github.com/mabd-dev/doc-gen-ai/internal/ollama"
)

type Pipeline struct {
	Ollama    *ollama.Client
	analyzer  analyzer
	generator generator
}

func NewPipeline(ollama *ollama.Client, verbose bool) *Pipeline {
	return &Pipeline{
		Ollama: ollama,
		analyzer: analyzer{
			MaxTries: 2,
			Client:   ollama,
			Verbose:  verbose,
		},
		generator: generator{
			MaxTries: 2,
			Client:   ollama,
			Verbose:  verbose,
		},
	}
}

func (p Pipeline) Analyze(code, prompt string) (string, error) {
	return p.analyzer.Analyze(code, prompt)
}

func (p Pipeline) GenerateDoc(
	analysis, signature, prompt string,
) (string, error) {
	return p.generator.Generate(analysis, signature, prompt)
}

// GetDocsOnly filter output from [GenerateDoc] function and get only the documentation part
func (p Pipeline) GetDocsOnly(docs string) (string, error) {
	return getDocsOnly(docs)
}

func (p Pipeline) PolishDoc(doc, prompt string) (string, error) {
	finalPrompt := strings.Replace(prompt, "{{KDOC}}", doc, 1)
	return p.Ollama.GenerateWithModel(finalPrompt, p.Ollama.DocPolishModel)
}
