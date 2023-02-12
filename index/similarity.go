package index

import (
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
	"math"
	"strings"
	"unicode"
)

// Tokenize splits a string into tokens
func Tokenize(text string) []string {
	tokens := strings.FieldsFunc(text, func(r rune) bool {
		return unicode.In(r, unicode.White_Space) ||
			unicode.In(r, unicode.Terminal_Punctuation) ||
			unicode.In(r, unicode.Sentence_Terminal) ||
			unicode.In(r, unicode.Quotation_Mark)
	})
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		panic(err)
	}
	result := make([]string, 0)
	for i := range tokens {
		wakatiTokens := t.Wakati(tokens[i])
		for j := range wakatiTokens {
			result = append(result, wakatiTokens[j])
		}
	}
	return result
}

// Vectorize creates a vector representation of a string
func Vectorize(text string) map[string]int {
	tokens := Tokenize(text)
	vector := make(map[string]int)
	for _, token := range tokens {
		vector[token]++
	}
	return vector
}

// Magnitude calculates the magnitude of a vector
func Magnitude(vector map[string]int) float64 {
	sum := 0.0
	for _, count := range vector {
		sum += float64(count * count)
	}
	return math.Sqrt(sum)
}

func CalculateSimilarity(s1 string, s2 string) float64 {
	vector1 := Vectorize(s1)
	vector2 := Vectorize(s2)

	dotProduct := 0.0
	for token, count := range vector1 {
		if vector2[token] > 0 {
			dotProduct += float64(count * vector2[token])
		}
	}

	magnitude1 := Magnitude(vector1)
	magnitude2 := Magnitude(vector2)

	return dotProduct / (magnitude1 * magnitude2)
}

func bytesToFloat64(ba []byte) []float64 {
	f := make([]float64, len(ba))
	for i := range f {
		f[i] = float64(ba[i])
	}
	return f
}
