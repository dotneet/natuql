package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	ApiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
	}
}

func (client *Client) request(endpoint string, requestBody map[string]interface{}) (map[string]interface{}, error) {
	if client.ApiKey == "" {
		return nil, fmt.Errorf("API Keyが指定されていません。")
	}

	// Marshal Request Body to JSON
	requestBodyBytes, _ := json.Marshal(requestBody)

	// Create a new HTTP Request
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBodyBytes))

	// Set API Key in the Request Header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.ApiKey))
	req.Header.Add("Content-Type", "application/json")

	// Send HTTP Request
	httpClient := &http.Client{}
	res, _ := httpClient.Do(req)

	// Read Response Body
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// Unmarshal Response Body to JSON
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	return response, nil
}
