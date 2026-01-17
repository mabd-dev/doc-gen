// Package pipeline
package pipeline

import (
	"github.com/mabd-dev/doc-gen-ai/internal/logger"
	"github.com/mabd-dev/doc-gen-ai/internal/ollama"
)

type Pipeline struct {
	logger    logger.Logger
	Ollama    *ollama.Client
	analyzer  analyzer
	generator generator
	polisher  polisher
}

func NewPipeline(
	ollama *ollama.Client,
	logger logger.Logger,
) *Pipeline {
	return &Pipeline{
		logger: logger,
		Ollama: ollama,
		analyzer: analyzer{
			MaxTries: 2,
			Client:   ollama,
			Logger:   logger,
		},
		generator: generator{
			MaxTries: 2,
			Client:   ollama,
			Logger:   logger,
		},
		polisher: polisher{
			MaxTries: 2,
			Client:   ollama,
			Logger:   logger,
		},
	}
}

func (p Pipeline) Analyze(code, prompt string) (string, error) {
	return p.analyzer.Analyze(code, prompt)
}

func (p Pipeline) GenerateDoc(
	analysis, prompt string,
) (string, error) {
	return p.generator.Generate(analysis, prompt)
}

func (p Pipeline) PolishDoc(docs, prompt string) (string, error) {
	return p.polisher.polish(docs, prompt)
}
