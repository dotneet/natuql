package openai

import (
	"fmt"
	"strings"
)

func (client *Client) CreateSql(context string, query string) (string, error) {
	prompt := fmt.Sprintf(`
Context information is below.
---------------------
DatabaseSchema: """
%s
"""

Example of Question and SQL: """
Natural:table から condition を抜き出して fileds 取得して。
SQL:SELECT ${fields} FROM ${table} WHERE ${condition}
"""
---------------------
Given the context information and not prior knowledge,
Natural:%s
SQL:
`, context, query)
	result, err := client.Complete(prompt)
	if err != nil {
		return "", err
	}
	fmt.Printf("Token usage: %f\n", result.TotalTokens)
	responseText := result.ResponseText
	start := "SQL:"
	idx := strings.Index(responseText, start)
	if idx >= 0 {
		sqlPos := idx + len(start)
		return responseText[sqlPos:], nil
	}
	idx = strings.Index(strings.ToLower(responseText), "select")
	if idx >= 0 {
		return responseText[idx:], nil
	}
	return "", fmt.Errorf("failed to create a sql.\nChatGPT Response: %s", responseText)
}
