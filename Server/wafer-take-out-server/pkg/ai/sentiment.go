package ai

import (
	"context"

	"github.com/tmc/langchaingo/llms/openai"
)

func GetDescriptionRanking(description string) (string, error) {

	key := ""

	key = OPENAI_API_KEY

	llm, err := openai.New(openai.WithToken(key))

	if err != nil {
		return description, err
	}

	template := BASE_PROMPT_TEMPLATE

	prompt := template + description

	response, err := llm.Call(context.Background(), prompt)

	if err != nil {
		return description, err
	}

	fullDescription := description + "\n-------AI评价-------\n" + response

	return fullDescription, nil
}
