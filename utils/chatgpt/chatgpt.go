package chatgpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

func GetChatGPTResponse(prompt string) (string, error) {
	var (
		chatGPTAPIKey = os.Getenv("OPENAI_KEY")
		chatGPTURL    = "https://api.openai.com/v1/chat/completions"
	)
	requestData := request{
		Model: "gpt-4-turbo-preview",
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

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", chatGPTURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+chatGPTAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch response. status code: " + resp.Status)
	}

	var result response
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	if len(result.Choices) < 1 {
		return "", errors.New("failed to fetch response")
	}
	return result.Choices[0].Message.Content, nil
}

type request struct {
	Model    string `json:"model"`
	Messages []struct {
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
