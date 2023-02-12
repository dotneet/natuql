package index

import "testing"
import "github.com/stretchr/testify/assert"

func TestChunks(t *testing.T) {
	s := &Schema{
		tables: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
	}
	chunks := s.Chunks(3)
	assert.Equal(t, 4, len(chunks))
	assert.Equal(t, []string{"a", "b", "c"}, chunks[0])
	assert.Equal(t, []string{"d", "e", "f"}, chunks[1])
	assert.Equal(t, []string{"g", "h", "i"}, chunks[2])
	assert.Equal(t, []string{"j"}, chunks[3])
}

func TestSchemaIndex_GetRelatedTables(t *testing.T) {
	index := &SchemaIndex{
		Tables: []string{"abc", "abd", "acd", "bcd", "bce", "bde", "bca"},
	}
	relatedTables := index.GetRelatedTables("abc", 5)
	assert.Equal(t, []string{"abc", "bca", "bcd", "bde", "acd"}, relatedTables)
}

func TestCalculateCosineSimilarity(t *testing.T) {
	similarity := CalculateSimilarity("利用者は誰？", "利用者は誰？")
	assert.Less(t, 0.999, similarity)

	similarity = CalculateSimilarity("君は誰？", "私達は誰？")
	assert.Greater(t, 1.0, similarity)

	similarity = CalculateSimilarity("a,b.c", "a,  b.c")
	assert.Less(t, 0.999, similarity)

	similarity = CalculateSimilarity("a b c", "c b a")
	assert.Less(t, 0.999, similarity)

	similarity = CalculateSimilarity("a b c", "a b d")
	assert.Greater(t, 1.0, similarity)
}
