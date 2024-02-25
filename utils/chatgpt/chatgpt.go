package chatgpt

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

var (
	chatGPTAPIKey = os.Getenv("OPENAI_KEY")
	chatGPTURL    = "https://api.openai.com/v1/chat/completions"
)

func GetChatGPTResponse(prompt string) (string, error) {
	client := resty.New()
	var result response
	request := request{
		Model: "gpt-4",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+chatGPTAPIKey).
		SetBody(request).SetResult(&result).
		Post(chatGPTURL)

	if err != nil {
		return "", err
	}

	fmt.Println("POST Response:", resp.Status())

	return result.Choices[0].Message.Content, nil
}

type request struct {
	Model     string `json:"model"`
	MaxTokens int    `json:"max_tokens"`
	Messages  []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}
type response struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
}
