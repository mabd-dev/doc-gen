package llm

type Client interface {
	GetBaseModel() string
	GetDocPolishModel() string

	Generate(prompt string) (string, error)
	GenerateWithModel(prompt, model string) (string, error)
}
