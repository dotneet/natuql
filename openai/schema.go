package openai

import (
	_ "embed"
	"fmt"
	"strings"
)

type Example struct {
	lang   string
	input  string
	output string
}

func (client *Client) AnnotateSchema(ddl string, lang string) (string, error) {
	example, err := client.createExample(lang)
	if err != nil {
		return "", err
	}
	return client.annotateSchemaWithExample(ddl, example)
}

func (client *Client) annotateSchemaWithExample(ddl string, example Example) (string, error) {
	promptSrc := `
Add comments to the DB schema.

Procedure:
 - Translate the table name in ${lang} and add it as the comment.
 - If the table name translated in ${lang} has similar words, add them too in ${lang}.
 - Translate the column name in ${lang} and add it as the comment.
 - If the column name translated in ${lang} has similar words, add them too in ${lang}.

Input Example: """
%s
"""

Output Example： """
%s
"""

Target: """
%s
"""
`
	replacedPromptSrc := strings.ReplaceAll(promptSrc, "${lang}", example.lang)
	prompt := fmt.Sprintf(replacedPromptSrc, example.input, example.output, ddl)
	result, err := client.Complete(prompt)
	if err != nil {
		return "", err
	}

	return result.ResponseText, nil
}

//go:embed japanese_example.txt
var japaneseExample string

//go:embed chinese_example.txt
var chineseExample string

//go:embed french_example.txt
var frenchExample string

func (client *Client) createExample(lang string) (Example, error) {
	exampleText := ""
	switch strings.ToLower(lang) {
	case "japanese":
		exampleText = japaneseExample
	case "chinese":
		exampleText = chineseExample
	case "french":
		exampleText = frenchExample
	default:
		return Example{}, fmt.Errorf("%s is unsupported", lang)
	}
	examples := strings.Split(exampleText, "=====")
	return Example{
		lang:   lang,
		input:  examples[0],
		output: examples[1],
	}, nil
}
