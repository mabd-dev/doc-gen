package llm

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type OllamaClient struct {
	BaseURL        string
	BaseModel      string
	DocPolishModel string
}

type ollamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ollamaResponse struct {
	Response string `json:"response"`
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
