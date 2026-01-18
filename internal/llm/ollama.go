package llm

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var (
	DefaultOllamaBaseURL    string = "http://localhost:11434"
	DefaultOllamaBaseModel  string = "qwen2.5-coder:7b"
	DefaultOllamaPolishDocs bool   = true
)

type OllamaClient struct {
	BaseURL        string
	BaseModel      string
	DocPolishModel string
	PolishDocs     bool
}

type ollamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ollamaResponse struct {
	Response string `json:"response"`
}

func (c OllamaClient) ShouldDoDocsPolishig() bool {
	return c.PolishDocs
}

func (c OllamaClient) GetBaseModel() string {
	return c.BaseModel
}

func (c OllamaClient) GetDocPolishModel() string {
	return c.DocPolishModel
}

func (c OllamaClient) Generate(prompt string) (string, error) {
	return c.GenerateWithModel(prompt, "")
}

func (c OllamaClient) GenerateWithModel(
	prompt string,
	model string,
) (string, error) {
	modelToUse := c.BaseModel
	if len(model) != 0 {
		modelToUse = model
	}

	reqBody, _ := json.Marshal(ollamaRequest{
		Model:  modelToUse,
		Prompt: prompt,
		Stream: false,
	})

	resp, err := http.Post(
		c.BaseURL+"/api/generate",
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result ollamaResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result.Response, err
}
