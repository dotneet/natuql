package index

import (
	"github.com/dotneet/natuql/path"
	"os"
)

func ReadIndex() (string, error) {
	indexFilePath, err := path.GetIndexFilePath()
	if err != nil {
		return "", err
	}

	bytes, err := os.ReadFile(indexFilePath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
