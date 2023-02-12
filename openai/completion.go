package openai

import (
	"fmt"
)

func (client *Client) Complete(prompt string) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("プロンプトが指定されていません。")
	}
	// API Endpoint
	endpoint := "https://api.openai.com/v1/completions"

	// Request Body
	requestBody := map[string]interface{}{
		"model":       "text-davinci-003",
		"prompt":      prompt,
		"max_tokens":  1000,
		"temperature": 0.5,
	}

	// Marshal Request Body to JSON
	response, err := client.request(endpoint, requestBody)
	if err != nil {
		return "", err
	}

	// Extract Generated Text from the Response
	generatedText := response["choices"].([]interface{})[0].(map[string]interface{})["text"].(string)

	return generatedText, nil
}
