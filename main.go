package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mabd-dev/doc-gen-ai/internal/llm"
	"github.com/mabd-dev/doc-gen-ai/internal/logger"
	"github.com/mabd-dev/doc-gen-ai/internal/pipeline"
	"github.com/mabd-dev/doc-gen-ai/internal/prompts"

	"golang.design/x/clipboard"
)

func main() {
	verboseFlag := flag.Bool("verbose", false, "Print `Debug`, `Warn` & `Error` to `stderr` ")
	vFlag := flag.Bool("v", false, "Print `Debug`, `Warn` & `Error` to `stderr` ")

	quietFlag := flag.Bool("quiet", false, "If false, print `Info` logs to stderr")
	qFlag := flag.Bool("q", false, "If false, print `Info` logs to stderr")

	providerFlag := flag.String("provider", "ollama", "Provider to use, options=(ollama, GROK)")
	pFlag := flag.String("p", "ollama", "Provider to use, options=(ollama, GROK)")

	readFromClipboard := flag.Bool("c", false, "Read code from clipboard")

	flag.Parse()

	logger := logger.Logger{
		Quiet:   *quietFlag || *qFlag,
		Verbose: *verboseFlag || *vFlag,
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

	selectedProvider := *providerFlag
	if pFlag != nil && len(strings.TrimSpace(*pFlag)) > 0 {
		selectedProvider = strings.TrimSpace(*pFlag)
	}

	client, err := llm.NewClient(selectedProvider)
	if err != nil {
		logger.LogError(err.Error())
		os.Exit(1)
	}

	pipeline := pipeline.NewPipeline(client, logger)

	// Step 1
	analysis, err := pipeline.Analyze(string(input), prompts.KotlinAnalyze)
	if err != nil {
		panic(err)
	}

	// Step 2
	docs, err := pipeline.GenerateDoc(analysis, prompts.KotlinKDoc)
	if err != nil {
		panic(err)
	}

	// Step 3
	docs, err = pipeline.PolishDoc(docs, prompts.KotlinKDocPolish)
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

func isValidKDoc(s string) bool {
	return strings.HasPrefix(strings.TrimSpace(s), "/**") &&
		strings.HasSuffix(strings.TrimSpace(s), "*/")
}
