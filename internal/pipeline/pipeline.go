// Package pipeline
package pipeline

import (
	"github.com/mabd-dev/doc-gen-ai/internal/ollama"
)

type Pipeline struct {
	Ollama    *ollama.Client
	analyzer  analyzer
	generator generator
	polisher  polisher
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
		polisher: polisher{
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

func (p Pipeline) PolishDoc(docs, prompt string) (string, error) {
	return p.polisher.polish(docs, prompt)
}
