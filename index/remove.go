package index

import (
	"fmt"
	"github.com/dotneet/natuql/path"
	"os"
)

func RemoveIndex() error {
	indexFilePath, err := path.GetIndexFilePath()
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
