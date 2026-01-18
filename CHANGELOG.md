# Changelog

## Unreleased

### ✨ Added
- Remote LLM support via OpenAI-compatible API
- Groq as first supported remote provider
- `--provider` flag to switch between `ollama` and `groq`

### Updates
- Add CLI flags to control `baseURL`, `baseModel` & `polishDocs` per provider


## [0.1.1] - 2026-01-17

### Updates
- Separate output streams: KDoc to stdout, logs to stderr
- Add `--quiet` flag to suppress info messages
- Add `--verbose` flag for debug/warn/error messages

## [0.1.0] - 2026-01-17

### ✨ Features

- Three-stage documentation pipeline: analyze → generate → polish
- Ollama integration for local LLM inference
- Clipboard support (read input, write output)
- Retry logic for generation and polish steps (max 2 attempts each)
- KDoc validation before output

### 🏗️ Architecture

- Modular struct-based design: DocAnalyzer, DocGenerator, DocPolisher
- Template-based prompt system for Kotlin KDoc generation

### 📚 Documentation

- Added function documentation throughout codebase
- Neovim integration guide with link to doc-gen-ai.nvim
