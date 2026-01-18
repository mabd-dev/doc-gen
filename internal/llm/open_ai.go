package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	defaultTemprature float32 = 0.2
	defaultTopP       float32 = 0.9
)

type OpenAIClient struct {
	BaseURL    string
	APIKey     string
	BaseModel  string
	PolishDocs bool
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAIRequest struct {
	Model            string    `json:"model"`
	Messages         []message `json:"messages"`
	Temperature      float32   `json:"temperature"`
	TopP             float32   `json:"top_p"`
	Stream           bool      `json:"stream"`
	MaxTokens        int       `json:"max_tokens,omitempty"`
	IncludeReasoning bool      `json:"include_reasoning"`
}

type openAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func (c OpenAIClient) ShouldDoDocsPolishig() bool {
	return c.PolishDocs
}

func (c OpenAIClient) GetBaseModel() string {
	return c.BaseModel
}

func (c OpenAIClient) GetDocPolishModel() string {
	return c.BaseModel
}

func (c OpenAIClient) Generate(prompt string) (string, error) {
	return c.GenerateWithModel(prompt, "")
}

func (c OpenAIClient) GenerateWithModel(prompt, model string) (string, error) {
	modelToUse := c.BaseModel
	if model != "" {
		modelToUse = model
	}

	reqBody, _ := json.Marshal(openAIRequest{
		Model: modelToUse,
		Messages: []message{
			{Role: "user", Content: prompt},
		},
		Temperature:      defaultTemprature,
		TopP:             defaultTopP,
		Stream:           false,
		MaxTokens:        4096,
		IncludeReasoning: false,
	})

	req, err := http.NewRequest("POST", c.BaseURL+"/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result openAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", fmt.Errorf("API error: %s", result.Error.Message)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no response from model")
	}

	// fmt.Println("______________")
	// for i, choise := range result.Choices {
	// 	fmt.Printf("result %i\n", i)
	// 	fmt.Println(choise.Message.Content)
	// }
	// fmt.Println("--------------")
	return result.Choices[0].Message.Content, nil
}
