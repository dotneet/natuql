package main

import (
	"github.com/dotneet/natuql/command"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func main() {
	var rootCmd *cobra.Command = nil
	rootCmd = &cobra.Command{
		Use:   "natuql",
		Short: "This tool allows you to use natural language as a query to a relational database.",
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd.Help()
		},
	}
	var indexCreateCmd = command.IndexCreateCmd()
	var indexRemoveCmd = command.IndexRemoveCmd()
	var queryCmd = command.QueryCommand()
	cobra.OnInitialize(func() {
		viper.SetDefault("apikey", os.Getenv("OPENAI_API_KEY"))
		viper.BindPFlag("apikey", rootCmd.PersistentFlags().Lookup("apikey"))
		viper.SetDefault("dbconn", os.Getenv("DATABASE_CONNECTION"))
		viper.BindPFlag("dbconn", rootCmd.PersistentFlags().Lookup("dbconn"))
	})
	rootCmd.PersistentFlags().String("apikey", "", "OpenAPI API Secret Key")
	rootCmd.PersistentFlags().String("dbconn", "", "Database Connection String")
	rootCmd.AddCommand(queryCmd)
	rootCmd.AddCommand(indexCreateCmd)
	rootCmd.AddCommand(indexRemoveCmd)
	rootCmd.Execute()
}
