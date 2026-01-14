package pipeline

import (
	"fmt"
	"strings"

	"github.com/mabd-dev/doc-gen-ai/internal/ollama"
)

type generator struct {
	MaxTries int
	Client   *ollama.Client
	Verbose  bool
}

func (g generator) Generate(
	analysis, signature, prompt string,
) (string, error) {

	for i := range g.MaxTries {

		if i == 0 {
			fmt.Println("Generating docs")
		} else {
			fmt.Printf("Generating docs, attempt %v/%v\n", i+1, g.MaxTries)
		}

		lastTry := i == g.MaxTries-1

		finalPrompt := strings.Replace(prompt, "{{ANALYSIS}}", analysis, 1)
		finalPrompt = strings.Replace(finalPrompt, "{{SIGNATURE}}", signature, 1)

		docs, err := g.Client.GenerateWithModel(finalPrompt, g.Client.BaseModel)

		if err != nil {
			if lastTry {
				return "", err
			} else {
				if g.Verbose {
					fmt.Println(err.Error())
				}
			}
		}

		docs, err = getDocsOnly(docs)

		if err != nil {
			if lastTry {
				return "", err
			} else {
				if g.Verbose {
					fmt.Println(err.Error())
				}
			}
		}

		if len(docs) != 0 {
			if g.Verbose {
				fmt.Println(docs)
			}
			return docs, nil
		}

	}

	err := fmt.Errorf("unable to generate documentation")
	return "", err
}
