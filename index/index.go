package index

import (
	"encoding/json"
	"os"
	"sort"
	"strings"
)

type Schema struct {
	tables []string
}

// json encodable schema index
type SchemaIndex struct {
	Tables []string `json:"tables"`
}

func (s *Schema) Chunks(size int32) [][]string {
	chunks := make([][]string, 0)
	for i := 0; i < len(s.tables); i += int(size) {
		end := i + int(size)
		if end > len(s.tables) {
			end = len(s.tables)
		}
		chunks = append(chunks, s.tables[i:end])
	}
	return chunks
}

func (index *SchemaIndex) WriteAsJson(path string) error {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(index)
}

type SortableItem struct {
	index      int
	similarity float64
	table      *string
}

func (schemaIndex *SchemaIndex) GetRelatedTablesString(query string, topK int) string {
	return strings.Join(schemaIndex.GetRelatedTables(query, topK), ";")
}

func (schemaIndex *SchemaIndex) GetRelatedTables(query string, topK int) []string {
	var items []SortableItem
	for i := 0; i < len(schemaIndex.Tables); i++ {
		table := schemaIndex.Tables[i]
		similarity := CalculateSimilarity(table, query)
		items = append(items, SortableItem{index: i, similarity: similarity, table: &table})
	}
	sort.Slice(items, func(i int, j int) bool {
		return items[i].similarity > items[j].similarity
	})

	items = items[0:topK]
	var result []string
	for i := 0; i < len(items); i++ {
		item := items[i]
		result = append(result, schemaIndex.Tables[item.index])
	}
	return result
}

func LoadSchemaIndexFromFile() (*SchemaIndex, error) {
	path, err := GetIndexFilePath()
	if err != nil {
		return nil, err
	}
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)
	var index SchemaIndex
	err = decoder.Decode(&index)
	if err != nil {
		return nil, err
	}
	return &index, nil
}
