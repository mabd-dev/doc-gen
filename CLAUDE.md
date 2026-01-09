# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go CLI tool that generates KDoc (Kotlin documentation) for Kotlin functions using local LLM models via Ollama. It takes Kotlin function code via stdin, analyzes it through a multi-stage pipeline, and outputs KDoc comments.

## Architecture

The codebase follows a three-stage pipeline architecture:

1. **Analysis Stage** (`main.go:51-65`): Analyzes the Kotlin function code to extract verifiable facts (return type, suspend status, error handling, side effects, etc.)
2. **Generation Stage** (`main.go:68-110`): Generates initial KDoc based on the analysis and function signature
3. **Polish Stage** (`main.go:112-145`): Refines the generated KDoc for clarity and consistency

### Key Components

- **`internal/pipeline`**: Core pipeline orchestration that coordinates the three stages (analyze, generate, polish)
- **`internal/ollama`**: HTTP client for communicating with local Ollama API at `localhost:11434`
- **`internal/prompts`**: Contains the LLM prompts used at each stage:
  - `KotlinAnalyze`: Extracts facts from code
  - `KotlinKDoc`: Generates KDoc from analysis
  - `KotlinKDocPolish`: Polishes generated documentation

### Pipeline Flow

The pipeline uses template-based prompt substitution (`{{FUNCTION}}`, `{{ANALYSIS}}`, `{{SIGNATURE}}`, `{{KDOC}}`). Each stage:
- Substitutes variables into prompt templates
- Calls Ollama API with the appropriate model
- Extracts KDoc from the response using regex (`/\*\*(.|[\r\n])*?\*/`)
- Validates the output before proceeding

### Retry Logic

Both the generation and polish stages implement retry logic (max 2 attempts each) to handle LLM generation failures. If retries are exhausted, the program panics with an error message.

### Models

The tool uses two models configured in `main.go:32-36`:
- `BaseModel`: "qwen-kdoc" - for analysis and initial generation
- `DocPolishModel`: "llama-kdoc:latest" - for polishing

## Development Commands

### Build the CLI
```bash
go build -o doc-gen
```

### Run the tool
```bash
# Pipe Kotlin function code to stdin
cat function.kt | go run main.go

# With verbose logging
cat function.kt | go run main.go -verbose
```

### Run tests
```bash
go test ./...
```

### Run tests for a specific package
```bash
go test ./internal/pipeline
go test ./internal/ollama
```

## Important Implementation Details

- The tool expects **Ollama to be running locally** at `http://localhost:11434`
- Input must be piped via **stdin** (the tool reads from `os.Stdin`)
- The function signature is extracted by searching for lines containing `"fun "` (`main.go:147-159`)
- KDoc validation requires the output to start with `/**` and end with `*/` (`main.go:161-164`)
- The `GetDocsOnly` function (`pipeline/pipeline.go:32-42`) uses regex to extract KDoc blocks from LLM responses that may contain additional text
- Error handling is strict: the program **panics** on failures rather than gracefully degrading

## Modifying Prompts

When editing prompts in `internal/prompts/`:
- Prompts use Go template-style placeholders: `{{FUNCTION}}`, `{{ANALYSIS}}`, `{{SIGNATURE}}`, `{{KDOC}}`
- Maintain the exact placeholder names as they're replaced via `strings.Replace` in the pipeline
- The `KotlinAnalyze` prompt defines a strict output format that must be preserved for downstream stages
