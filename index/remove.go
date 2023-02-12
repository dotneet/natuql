package index

import (
	"fmt"
	"os"
)

func RemoveIndex() error {
	indexFilePath, err := GetIndexFilePath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// check if index file exists
	_, err = os.Stat(indexFilePath)
	if err != nil {
		return nil
	}
	return os.Remove(indexFilePath)
}
