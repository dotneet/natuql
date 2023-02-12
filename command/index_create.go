package command

import (
	"fmt"
	"github.com/dotneet/natuql/index"
	"github.com/dotneet/natuql/openai"
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
			connStr := viper.GetString("dbconn")
			apiKey := viper.GetString("apikey")
			client := openai.NewClient(apiKey)
			err := index.CreateIndex(client, driverName, connStr)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error: %v", err)
				return
			}
			path, err := index.GetIndexFilePath()
			if err != nil {
				fmt.Fprintln(os.Stderr, "error: %v", err)
				return
			}
			fmt.Println("index has been created.")
			fmt.Println("location: %s", path)
		},
	}
}
