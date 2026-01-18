package pipeline

import (
	"fmt"
	"strings"

	"github.com/mabd-dev/doc-gen-ai/internal/llm"
	"github.com/mabd-dev/doc-gen-ai/internal/logger"
)

type polisher struct {
	MaxTries int
	Client   llm.Client
	Logger   logger.Logger
}

func (p polisher) polish(docs, prompt string) (string, error) {
	for i := range p.MaxTries {
		if i == 0 {
			p.Logger.LogInfo("Polishing docs...")
		} else {
			p.Logger.LogInfo("Polishing docs, attempts %v/%v\n", i, p.MaxTries)
		}

		lastTry := i == p.MaxTries-1

		finalPrompt := strings.Replace(prompt, "{{KDOC}}", docs, 1)
		polishedDocs, err := p.Client.GenerateWithModel(finalPrompt, p.Client.GetDocPolishModel())

		if err != nil {
			if lastTry {
				return "", err
			}
			p.Logger.LogError(err.Error())
		}

		polishedDocs, err = getDocsOnly(polishedDocs)

		if err != nil {
			if lastTry {
				return "", err
			}
			p.Logger.LogError(err.Error())
		}

		if len(polishedDocs) != 0 {
			p.Logger.LogDebug(polishedDocs)
			return polishedDocs, nil
		}

	}

	err := fmt.Errorf("failed to polish docs")
	return "", err
}
