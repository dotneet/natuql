package index

import (
	"fmt"
	"os"
	"os/user"
)

var indexFileName = "natuql.index"

func GetIndexFilePath() (string, error) {
	configDirPath, err := getConfigDirectoryPath()
	if err != nil {
		return "", err
	}
	return configDirPath + "/" + indexFileName, nil
}

// get config directory and create if not exists
func getConfigDirectoryPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ホームディレクトリのパスを取得する
	homeDir := usr.HomeDir
	configDirPath := homeDir + "/.config/natuql"
	// check if config directory exists
	_, err = os.Stat(configDirPath)
	if err != nil {
		// create config directory
		err := os.MkdirAll(configDirPath, 0755)
		if err != nil {
			return "", err
		}
	}
	return configDirPath, nil
}
