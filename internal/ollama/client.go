package ollama

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	BaseURL        string
	BaseModel      string
	DocPolishModel string
}

type request struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type response struct {
	Response string `json:"response"`
}

func (c *Client) Generate(prompt string) (string, error) {
	return c.GenerateWithModel(prompt, "")
}

func (c *Client) GenerateWithModel(
	prompt string,
	model string,
) (string, error) {
	modelToUse := c.BaseModel
	if len(model) != 0 {
		modelToUse = model
	}

	reqBody, _ := json.Marshal(request{
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

	var result response
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result.Response, err
}
