package openai

import (
	"fmt"
)

type CompletionResult struct {
	ResponseText    string
	TotalTokens     float64
	PromptTokens    float64
	CompletionToken float64
}

func (client *Client) Complete(prompt string) (CompletionResult, error) {
	if prompt == "" {
		return CompletionResult{}, fmt.Errorf("prompt is not specified")
	}
	// API Endpoint
	endpoint := "https://api.openai.com/v1/completions"

	// Request Body
	requestBody := map[string]interface{}{
		"model":       client.Model,
		"prompt":      prompt,
		"max_tokens":  client.MaxToken,
		"temperature": client.Temperature,
	}

	// Marshal Request Body to JSON
	response, err := client.request(endpoint, requestBody)
	if err != nil {
		return CompletionResult{}, err
	}

	if response["error"] != nil {
		errorField := response["error"].(map[string]interface{})
		if errorField != nil {
			return CompletionResult{}, fmt.Errorf("ChatGPT API Error: %s", errorField["message"].(string))
		}
	}

	// Extract Generated Text from the Response
	generatedText := response["choices"].([]interface{})[0].(map[string]interface{})["text"].(string)
	usage := response["usage"].(interface{}).(map[string]interface{})
	promptTokens := usage["prompt_tokens"].(float64)
	completionTokens := usage["completion_tokens"].(float64)
	totalTokens := usage["total_tokens"].(float64)

	return CompletionResult{
		ResponseText:    generatedText,
		TotalTokens:     totalTokens,
		PromptTokens:    promptTokens,
		CompletionToken: completionTokens,
	}, nil
}
