# KDoc Generator

Automatically generate high-quality KDoc documentation for your Kotlin functions using local AI models.

## Why This Exists

Documentation is critical for maintainable code, yet it's often the first thing to fall behind in fast-paced development. Manual documentation is time-consuming, inconsistent, and rarely reflects the actual behavior of the code. This tool bridges that gap.

## The Documentation Problem

Modern software teams face a persistent challenge:

- **Time Pressure**: Writing comprehensive documentation takes time developers don't have
- **Inconsistency**: Different developers document in different styles and levels of detail
- **Staleness**: Documentation quickly becomes outdated as code evolves
- **Cognitive Load**: Context-switching between coding and writing documentation breaks flow
- **Onboarding Friction**: New team members struggle to understand undocumented or poorly documented code

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
cat your-function.kt | ./doc-gen

# With verbose output to see the analysis stages
cat your-function.kt | ./doc-gen -verbose
```

## Use Cases

- **Legacy Code Modernization**: Quickly add documentation to undocumented codebases
- **API Documentation**: Generate consistent documentation for public APIs
- **Code Reviews**: Ensure all new functions have proper KDoc before merging
- **Onboarding**: Help new developers understand existing code through generated docs
- **CI/CD Integration**: Automatically verify or generate documentation in your build pipeline

## Requirements

- Go 1.x or later
- Ollama running locally with compatible models

## The Bottom Line

Good documentation shouldn't be a luxury. This tool makes it the default. Spend your time writing code that matters, and let AI handle the documentation that makes it understandable.

---

**Built with**: Go | Ollama | Local AI Models
