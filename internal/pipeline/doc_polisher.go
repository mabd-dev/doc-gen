package pipeline

import (
	"fmt"
	"strings"

	"github.com/mabd-dev/doc-gen-ai/internal/ollama"
)

type polisher struct {
	MaxTries int
	Client   *ollama.Client
	Verbose  bool
}

func (p polisher) polish(docs, prompt string) (string, error) {
	for i := range p.MaxTries {
		if i == 0 {
			fmt.Println("Polishing docs")
		} else {
			fmt.Printf("Polishing docs, attempts %v/%v\n", i, p.MaxTries)
		}

		lastTry := i == p.MaxTries-1

		finalPrompt := strings.Replace(prompt, "{{KDOC}}", docs, 1)
		polishedDocs, err := p.Client.GenerateWithModel(finalPrompt, p.Client.DocPolishModel)

		if err != nil {
			if lastTry {
				return "", err
			}
			if p.Verbose {
				fmt.Println(err.Error())
			}
		}

		polishedDocs, err = getDocsOnly(polishedDocs)

		if err != nil {
			if lastTry {
				return "", err
			}
			if p.Verbose {
				fmt.Println(err.Error())
			}
		}

		if len(polishedDocs) != 0 {
			if p.Verbose {
				fmt.Println(polishedDocs)
			}
			return polishedDocs, nil
		}

	}

	err := fmt.Errorf("failed to polish docs")
	return "", err
}
