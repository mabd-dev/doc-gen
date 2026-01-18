package pipeline

import (
	"fmt"
	"strings"

	"github.com/mabd-dev/doc-gen-ai/internal/llm"
	"github.com/mabd-dev/doc-gen-ai/internal/logger"
)

type generator struct {
	MaxTries int
	Client   llm.Client
	Logger   logger.Logger
}

func (g generator) Generate(
	analysis, prompt string,
) (string, error) {

	for i := range g.MaxTries {

		if i == 0 {
			g.Logger.LogInfo("Generating docs...")
		} else {
			g.Logger.LogInfo("Generating docs, attempt %v/%v\n", i+1, g.MaxTries)
		}

		lastTry := i == g.MaxTries-1

		finalPrompt := strings.Replace(prompt, "{{ANALYSIS}}", analysis, 1)

		docs, err := g.Client.GenerateWithModel(finalPrompt, g.Client.GetBaseModel())

		if err != nil {
			if lastTry {
				return "", err
			} else {
				g.Logger.LogError(err.Error())
			}
		}

		docs, err = getDocsOnly(docs)

		if err != nil {
			if lastTry {
				return "", err
			} else {
				g.Logger.LogError(err.Error())
			}
		}

		if len(docs) != 0 {
			g.Logger.LogDebug(docs)
			return docs, nil
		}

	}

	err := fmt.Errorf("unable to generate documentation")
	return "", err
}
