package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mabd-dev/doc-gen-ai/internal/ollama"
	"github.com/mabd-dev/doc-gen-ai/internal/pipeline"
	"github.com/mabd-dev/doc-gen-ai/internal/prompts"
)

func main() {
	verbose := flag.Bool("verbose", false, "Print verbose logs")

	flag.Parse()

	input, _ := io.ReadAll(os.Stdin)

	functionSignature := extractSignature(string(input))
	if len(functionSignature) == 0 {
		fmt.Println("Could not find function signature")
		os.Exit(1)
	}

	if *verbose {
		fmt.Println(functionSignature)
	}

	client := &ollama.Client{
		BaseURL:        "http://localhost:11434",
		BaseModel:      "qwen-kdoc",
		DocPolishModel: "llama-kdoc:latest",
	}

	p := pipeline.NewPipeline(client, *verbose)

	// Step 1
	analysis, err := p.Analyze(string(input), prompts.KotlinAnalyze)
	if err != nil {
		panic(err)
	}

	// Step 2
	docs, err := p.GenerateDoc(analysis, functionSignature, prompts.KotlinKDoc)
	if err != nil {
		panic(err)
	}

	// Step 3
	docs = polishDocs(*p, docs)

	if !isValidKDoc(docs) {
		panic(fmt.Errorf("failed to generate docs, result is not valid, result=\n%v", docs))
	}

	fmt.Println(docs)
}

// polishDocs takes a docs and improve clarity, writing lanuage etc...
//
// Parameters:
//   - docs: valid kdoc string
//
// Output:
//   - pure valid kdoc string
func polishDocs(p pipeline.Pipeline, docs string) string {
	fmt.Println("Polishing docs...")

	tries := 2
	for i := range tries {
		lastTry := i == tries-1

		polishedKDoc, err := p.PolishDoc(docs, prompts.KotlinKDocPolish)
		if err != nil {
			if lastTry {
				panic(fmt.Errorf("failed to get polished docs, error=%v", err.Error()))
			}
		}

		cleanedDocs, err := p.GetDocsOnly(polishedKDoc)
		cleanedDocs = strings.TrimSpace(cleanedDocs)

		if err == nil {
			if len(cleanedDocs) != 0 {
				return cleanedDocs
			}
		} else {
			if lastTry {
				fmt.Println(polishedKDoc)
				panic(fmt.Errorf("failed to extract kdoc from polished docs, error=%v", err.Error()))
			}
		}

		fmt.Println("Failed to generate polished docs!")
		fmt.Printf("Polishing docs: Try %v:%v\n", i+1, tries)
	}

	panic(fmt.Errorf("failed to generate polished doc, docs=\n%v", docs))
}

func extractSignature(function string) string {
	lines := strings.SplitSeq(function, "\n")
	for line := range lines {
		if strings.Contains(line, "fun ") {
			signature := strings.Split(line, "{")
			if len(signature) == 0 {
				return ""
			}
			return signature[0]
		}
	}
	return ""
}

func isValidKDoc(s string) bool {
	return strings.HasPrefix(strings.TrimSpace(s), "/**") &&
		strings.HasSuffix(strings.TrimSpace(s), "*/")
}
