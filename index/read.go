package index

import (
	"os"
)

func ReadIndex() (string, error) {
	indexFilePath, err := getIndexFilePath()
	if err != nil {
		return "", err
	}

	bytes, err := os.ReadFile(indexFilePath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
