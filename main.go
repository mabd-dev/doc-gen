package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mabd-dev/doc-gen/internal/ollama"
	"github.com/mabd-dev/doc-gen/internal/pipeline"
	"github.com/mabd-dev/doc-gen/internal/prompts"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)

	functionSignature := extractSignature(string(input))
	if len(functionSignature) == 0 {
		fmt.Println("Could not find function signature")
		os.Exit(1)
	}

	fmt.Println(functionSignature)

	client := &ollama.Client{
		BaseURL:        "http://localhost:11434",
		BaseModel:      "qwen-kdoc",
		DocPolishModel: "llama-kdoc:latest",
	}

	p := &pipeline.Pipeline{Ollama: client}

	fmt.Println("Analyzing code...")
	analysis, err := p.Analyze(string(input), prompts.KotlinAnalyze)
	if err != nil {
		panic(err)
	}
	fmt.Println("======= Analysis ========")
	fmt.Println(analysis)
	fmt.Println("=============")

	fmt.Println("Generating docs...")
	docs, err := p.GenerateDoc(analysis, functionSignature, prompts.KotlinKDoc)
	if err != nil {
		panic(err)
	}

	cleanedDocs, err := p.GetDocsOnly(docs)
	if err != nil {
		panic(err)
	}

	fmt.Println("====== KDOC =======")
	fmt.Println(cleanedDocs)
	fmt.Println("=============")

	polishedKDoc, err := p.PolishDoc(docs, prompts.KotlinKDocPolish)
	if err != nil {
		panic(err)
	}

	cleanedDocs, err = p.GetDocsOnly(polishedKDoc)
	if err != nil {
		panic(err)
	}

	fmt.Println("====== POLISHED KDOC =======")
	fmt.Println(cleanedDocs)
	fmt.Println("=============")
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
