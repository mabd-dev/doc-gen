package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mabd-dev/doc-gen-ai/internal/logger"
	"github.com/mabd-dev/doc-gen-ai/internal/ollama"
	"github.com/mabd-dev/doc-gen-ai/internal/pipeline"
	"github.com/mabd-dev/doc-gen-ai/internal/prompts"

	"golang.design/x/clipboard"
)

func main() {
	verbose := flag.Bool("verbose", false, "Print `Debug`, `Warn` & `Error` to `stderr` ")
	v := flag.Bool("v", false, "Print `Debug`, `Warn` & `Error` to `stderr` ")

	quiet := flag.Bool("quiet", false, "If false, print `Info` logs to stderr")
	q := flag.Bool("q", false, "If false, print `Info` logs to stderr")

	readFromClipboard := flag.Bool("c", false, "Read code from clipboard")

	flag.Parse()

	logger := logger.Logger{
		Quiet:   *quiet || *q,
		Verbose: *verbose || *v,
	}

	errorLoadingClipboard := clipboard.Init() != nil
	if errorLoadingClipboard {
		logger.LogError("WARNING: error initializing clipboard manager")
	}

	var input []byte
	if *readFromClipboard {
		input = clipboard.Read(clipboard.FmtText)
	} else {
		input, _ = io.ReadAll(os.Stdin)
	}

	functionSignature := extractSignature(string(input))
	if len(functionSignature) == 0 {
		logger.LogError("Could not find function signature")
		os.Exit(1)
	}

	logger.LogDebug(functionSignature)

	client := &ollama.Client{
		BaseURL:        "http://localhost:11434",
		BaseModel:      "qwen-kdoc",
		DocPolishModel: "llama-kdoc:latest",
	}

	p := pipeline.NewPipeline(client, logger)

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
	docs, err = p.PolishDoc(docs, prompts.KotlinKDocPolish)
	if err != nil {
		panic(err)
	}

	if !isValidKDoc(docs) {
		panic(fmt.Errorf("failed to generate docs, result is not valid, result=\n%v", docs))
	}

	fmt.Println(docs)
	if !errorLoadingClipboard {
		clipboard.Write(clipboard.FmtText, []byte(docs))
		logger.LogDebug("docs written to clipboard")
	}
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
