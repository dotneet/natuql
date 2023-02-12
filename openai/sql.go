package openai

import (
	"fmt"
	"strings"
)

func (client *Client) CreateSql(context string, query string) (string, error) {
	prompt := fmt.Sprintf(`
Context information is below. 
---------------------
データベーステーブルの定義: """
%s
"""
Q:TからQのものについてXを取得して。
SQL:select {X} from {Q} where {T}
"""
---------------------
Given the context information and not prior knowledge,
Q:%s
`, context, query)
	response, err := client.Complete(prompt)
	if err != nil {
		return "", err
	}
	start := "SQL:"
	idx := strings.Index(response, start)
	if idx == -1 {
		return "", err
	}
	sqlPos := idx + len(start)
	return response[sqlPos:], nil
}
