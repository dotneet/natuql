package index

import (
	"os"
)

func ReadIndex() (string, error) {
	indexFilePath, err := GetIndexFilePath()
	if err != nil {
		return "", err
	}

	bytes, err := os.ReadFile(indexFilePath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
