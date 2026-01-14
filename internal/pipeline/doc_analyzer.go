package pipeline

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mabd-dev/doc-gen-ai/internal/ollama"
)

type Visibility string

const (
	Public    Visibility = "public"
	Internal  Visibility = "internal"
	Protected Visibility = "protected"
	Private   Visibility = "private"
)

type Nullability bool

type Confidence string

const (
	HighConfidence   Confidence = "high"
	MediumConfidence Confidence = "medium"
	LowConfidence    Confidence = "low"
)

type Parameter struct {
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Nullable   Nullability `json:"nullable,omitempty"`
	Default    *string     `json:"default,omitempty"`
	Usage      string      `json:"usage"`
	Confidence Confidence  `json:"confidence"`
}

type Receiver struct {
	Type     string      `json:"type"`
	Nullable Nullability `json:"nullable,omitempty"`
}

type Return struct {
	Type       string      `json:"type"`
	Nullable   Nullability `json:"nullable,omitempty"`
	Semantics  string      `json:"semantics"`
	Confidence Confidence  `json:"confidence"`
	UsesResult bool        `json:"uses_result"`
}

type SideEffect struct {
	Description string     `json:"description"`
	Confidence  Confidence `json:"confidence"`
}

type Throw struct {
	Type       string     `json:"type"`
	Reason     string     `json:"reason"`
	Origin     string     `json:"origin"`
	Confidence Confidence `json:"confidence"`
}

type ControlFlow struct {
	EarlyReturns bool `json:"early_returns"`
	UsesTryCatch bool `json:"uses_try_catch"`
}

type ConfidenceWrapper struct {
	Overall Confidence `json:"overall"`
	Notes   string     `json:"notes,omitempty"`
}

type FunctionAnalysis struct {
	SummaryHint string            `json:"summary_hint"`
	Visibility  Visibility        `json:"visibility"`
	Modifiers   []string          `json:"modifiers"`
	Receiver    *Receiver         `json:"receiver,omitempty"`
	Parameters  []Parameter       `json:"parameters"`
	Return      Return            `json:"return"`
	SideEffects []SideEffect      `json:"side_effects"`
	Throws      []Throw           `json:"throws"`
	ControlFlow ControlFlow       `json:"control_flow"`
	Confidence  ConfidenceWrapper `json:"confidence"`
}

type analyzer struct {
	MaxTries int
	Client   *ollama.Client
	Verbose  bool
}

func (a analyzer) Analyze(code, prompt string) (string, error) {
	finalPrompt := strings.Replace(prompt, "{{FUNCTION}}", code, 1)

	var err error
	analysis := ""

	fmt.Println("Analyzing code...")

	for i := range a.MaxTries {
		if i != 0 {
			fmt.Printf("Analyzing code: Attempt %v/%v\n", i, a.MaxTries)
		}
		lastTry := i == a.MaxTries-1

		analysis, err = a.Client.GenerateWithModel(finalPrompt, a.Client.BaseModel)
		if err != nil {
			if lastTry {
				return "", err
			} else {
				continue
			}
		}

		err = a.validate(analysis)
		if err != nil {
			if lastTry {
				return "", nil
			}
		} else {
			break
		}
	}

	if len(analysis) == 0 {
		return "", fmt.Errorf("failed to analyze code, tried %d times", a.MaxTries)
	}

	if a.Verbose {
		fmt.Println(analysis)
	}

	return analysis, nil

}

func (a analyzer) validate(output string) error {
	var analysis FunctionAnalysis

	err := json.Unmarshal([]byte(output), &analysis)
	if err != nil {
		return err
	}

	// Validating summary hint
	if len(strings.TrimSpace(analysis.SummaryHint)) == 0 {
		return fmt.Errorf("summary hint cannot be empyt")
	}

	// Validating visibility
	switch analysis.Visibility {
	case Public, Internal, Protected, Private:
		break
	default:
		return fmt.Errorf("unknown visibility=%v", string(analysis.Visibility))
	}

	// TODO: add more validation on analysis

	return nil
}
