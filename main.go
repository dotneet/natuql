package main

import (
	"fmt"
	"github.com/dotneet/natuql/command"
	"github.com/dotneet/natuql/path"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
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
	viper.SetDefault("language", detectLanguage())
	cobra.OnInitialize(func() {
		viper.SetDefault("apikey", os.Getenv("OPENAI_API_KEY"))
		viper.BindPFlag("apikey", rootCmd.PersistentFlags().Lookup("apikey"))
		viper.SetDefault("dbconn", os.Getenv("DATABASE_CONNECTION"))
		viper.BindPFlag("dbconn", rootCmd.PersistentFlags().Lookup("dbconn"))
		viper.SetDefault("context-tables-count", 8)
	})
	rootCmd.PersistentFlags().String("apikey", "", "OpenAPI API Secret Key")
	rootCmd.PersistentFlags().String("dbconn", "", "Database Connection String")
	rootCmd.AddCommand(queryCmd)
	rootCmd.AddCommand(indexCreateCmd)
	rootCmd.AddCommand(indexRemoveCmd)

	viper.SetConfigName("config")
	configDirPath, err := path.GetConfigDirectoryPath()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	viper.AddConfigPath(configDirPath)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Fprint(os.Stderr, err)
			return
		}
	}
	rootCmd.Execute()
}

func detectLanguage() string {
	langCode := os.Getenv("LANG")
	if langCode == "" {
		langCode = os.Getenv("LC_CTYPE")
	}
	langParts := strings.Split(langCode, "_")
	country := langParts[0]
	switch strings.ToLower(country) {
	case "ja":
		return "Japanese"
	case "en":
		return "English"
	case "fr":
		return "French"
	case "zh":
		return "Chinese"
	default:
		return "English"
	}
}
