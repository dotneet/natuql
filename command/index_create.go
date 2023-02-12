package command

import (
	"fmt"
	"github.com/dotneet/natuql/index"
	"github.com/dotneet/natuql/openai"
	"github.com/dotneet/natuql/path"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func IndexCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "index-create",
		Short:   "create an index from the database.",
		Example: "natuql index-create",
		Run: func(cmd *cobra.Command, args []string) {
			driverName := "mysql"
			lang := viper.GetString("language")
			connStr := viper.GetString("dbconn")
			apiKey := viper.GetString("apikey")
			client := openai.NewClient(apiKey)
			path, err := path.GetIndexFilePath()
			if err != nil {
				fmt.Fprintln(os.Stderr, "error: %v", err)
				return
			}
			// check if index file already exists
			_, err = os.Stat(path)
			if err == nil {
				// read line from stdin
				fmt.Printf("index file already exists. overwrite? [y/N]: ")
				var input string
				fmt.Scanln(&input)
				if input != "y" {
					fmt.Println("canceled.")
					return
				}
			}

			err = index.CreateIndex(client, driverName, connStr, lang)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error: %v", err)
				return
			}
			fmt.Println("index has been created.")
			fmt.Printf("location: %s\n", path)
		},
	}
}
