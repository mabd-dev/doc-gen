# Changelog


## UnReleased

### Updates
- Direct all logs to `stderr` and only kdoc output on `stdout`
- Added `quiet` flag to silent all info logs
- Make `verbose` flag control `debug`, `warn` and `error` logs

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
