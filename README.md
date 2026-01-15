# KDoc Generator

Automatically generate high-quality KDoc documentation for your Kotlin functions using local AI models.

## Why This Exists

Documentation is critical for maintainable code, yet it's often the first thing to fall behind in fast-paced development. Manual documentation is time-consuming, inconsistent, and rarely reflects the actual behavior of the code. This tool bridges that gap.

## How This Tool Helps

This CLI tool analyzes your Kotlin functions and generates accurate, consistent KDoc comments automatically. Instead of spending hours writing documentation, you get:

### 1. **Instant Documentation Coverage**
Generate comprehensive KDoc for any Kotlin function in seconds. Transform undocumented legacy code into well-documented APIs without manual effort.

### 2. **Consistent Quality and Style**
Every function gets documented using the same analytical approach and formatting standards. No more variation between team members or documentation styles.

### 3. **Accurate Technical Details**
The tool uses a three-stage analysis pipeline that:
- Extracts verifiable facts (return types, suspend status, error handling)
- Generates documentation based on actual code behavior
- Polishes the output for clarity and readability

### 4. **Developer Productivity**
Stay in your coding flow. Pipe a function to the tool, get documentation back. No context switching, no breaking concentration to write prose.

### 5. **Better Code Understanding**
The tool analyzes side effects, error handling, and suspension behavior—details that are easy to miss in manual documentation but critical for correct API usage.

### 6. **Privacy-First Approach**
Uses local LLM models via Ollama. Your code never leaves your machine, making it suitable for proprietary and sensitive codebases.

## Quick Start

```bash
# Generate documentation for a Kotlin function
cat your-function.kt | ./doc-gen-ai

# With verbose output to see the analysis stages
cat your-function.kt | ./doc-gen-ai -verbose

# Read kotlin function from clipboard
./doc-gen-ai -c
```

## Use Cases

- **Legacy Code Modernization**: Quickly add documentation to undocumented codebases
- **API Documentation**: Generate consistent documentation for public APIs
- **Code Reviews**: Ensure all new functions have proper KDoc before merging
- **Onboarding**: Help new developers understand existing code through generated docs
- **CI/CD Integration**: Automatically verify or generate documentation in your build pipeline

## Requirements

- Go 1.24 or later
- Ollama running locally with compatible models

## Local Setup

This tool has been tested and optimized on the following configuration:

**Hardware**
- MacBook Pro M3 Pro with 18GB RAM

**AI Models** (via Ollama)
- **qwen2.5-coder:7b** - Used for code analysis and documentation generation
- **llama3.2:3b** - Used for documentation polishing and refinement

## Neovim Setup

- User Command

Add this to `init.lua` or `init.vim`
```lua
vim.api.nvim_create_user_command('DocGen', function(opts)
    local range = opts.line1 .. ',' .. opts.line2
    local command = 'w !doc-gen-ai ' .. range
    vim.cmd(command)
end, { range = true })
```

run it using
```sh
:'<,'>DocGen
```

- Keymap

```lua
vim.keymap.set("v", "<leader>d", ":DocGen<CR>", { desc = "Generate KODoc for selected kotlin function" })
```

## Pipeline

1. **Analysis Stage**: Analyzes the Kotlin function code to extract verifiable facts (return type, suspend status, error handling, side effects, etc.)
2. **Generation Stage**: Generates initial KDoc based on the analysis and function signature
3. **Polish Stage**: Refines the generated KDoc for clarity and consistency


