package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mabd-dev/doc-gen/internal/ollama"
	"github.com/mabd-dev/doc-gen/internal/pipeline"
	"github.com/mabd-dev/doc-gen/internal/prompts"
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

	p := &pipeline.Pipeline{Ollama: client}

	analysis := analyizeFunction(*p, string(input), *verbose)
	docs := generateDocs(*p, analysis, functionSignature, *verbose)
	docs = polishDocs(*p, docs)

	if !isValidKDoc(docs) {
		panic(fmt.Errorf("failed to generate docs, result is not valid, result=\n%v", docs))
	}

	fmt.Println(docs)
}

func analyizeFunction(
	p pipeline.Pipeline,
	code string,
	verbose bool,
) string {
	fmt.Println("Analyzing code...")
	analysis, err := p.Analyze(code, prompts.KotlinAnalyze)
	if err != nil {
		panic(err)
	}

	if verbose {
		fmt.Println(analysis)
	}
	return analysis
}

func generateDocs(
	p pipeline.Pipeline,
	analysis, functionSignature string,
	verbose bool,
) string {
	cleanedDocs := ""

	// Try max of 2 times to generate docs, if failed -> panic
	for i := range 2 {
		fmt.Println("Generating docs...")
		docs, err := p.GenerateDoc(analysis, functionSignature, prompts.KotlinKDoc)
		if err != nil {
			panic(err)
		}

		docs, err = p.GetDocsOnly(docs)
		if err != nil {
			panic(err)
		}

		if verbose {
			fmt.Println("generated docs:")
			fmt.Println(docs)
		}

		if len(docs) != 0 {
			cleanedDocs = docs
			break
		}
		fmt.Println("Failed to generate docs!")
		fmt.Printf("Try %v: ", i+1)
	}

	if len(cleanedDocs) == 0 {
		panic("unable to generate documentation")
	}

	if verbose {
		fmt.Println(cleanedDocs)
	}

	return cleanedDocs
}

func polishDocs(p pipeline.Pipeline, docs string) string {
	fmt.Println("Polishing docs...")
	polishedKDoc, err := p.PolishDoc(docs, prompts.KotlinKDocPolish)
	if err != nil {
		panic(fmt.Errorf("failed to get polished docs, error=%v", err.Error()))
	}

	cleanedDocs, err := p.GetDocsOnly(polishedKDoc)
	if err != nil {
		fmt.Println(polishedKDoc)
		panic(fmt.Errorf("failed to extract kdoc from polished docs, error=%v", err.Error()))
	}
	return cleanedDocs
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
