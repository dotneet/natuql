package main

import (
	"database/sql"
	"fmt"
	"github.com/dotneet/natuql/openai"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"

	"github.com/dotneet/natuql/index"
	_ "github.com/go-sql-driver/mysql"
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
	var indexCmd = &cobra.Command{
		Use:     "create-index",
		Short:   "create index from database.",
		Example: "natuql create-index",
		Run: func(cmd *cobra.Command, args []string) {
			driverName := "mysql"
			connStr := viper.GetString("dbconn")
			fmt.Println(connStr)
			apiKey := viper.GetString("apikey")
			client := openai.NewClient(apiKey)
			err := index.CreateIndex(client, driverName, connStr)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error: %v", err)
			}
		},
	}
	var queryCmd = &cobra.Command{
		Use:     "query",
		Short:   "query to ChatGPT",
		Example: "natuql query \"2022年に一番売れた商品とその合計価格を出力して。\"",
		Run: func(cmd *cobra.Command, args []string) {
			apiKey := viper.GetString("apikey")
			client := openai.NewClient(apiKey)
			query := strings.Join(args, " ")
			indexStr, err := index.ReadIndex()
			if err != nil {
				// print error to stderr
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				return
			}
			sql, err := client.CreateSql(indexStr, query)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				return
			}
			fmt.Printf("SQL: %s\n", sql)
			connStr := viper.GetString("dbconn")
			rows, err := executeSql("mysql", connStr, sql)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				return
			}
			fmt.Println(strings.Join(rows, "\n"))
		},
	}
	cobra.OnInitialize(func() {
		viper.SetDefault("apikey", os.Getenv("OPENAI_API_KEY"))
		viper.BindPFlag("apikey", rootCmd.PersistentFlags().Lookup("apikey"))
		viper.SetDefault("dbconn", os.Getenv("DATABASE_CONNECTION"))
		viper.BindPFlag("dbconn", rootCmd.PersistentFlags().Lookup("dbconn"))
	})

	rootCmd.PersistentFlags().String("apikey", "", "OpenAPI API Secret Key")
	rootCmd.PersistentFlags().String("dbconn", "", "Database Connection String")
	rootCmd.AddCommand(queryCmd)
	rootCmd.AddCommand(indexCmd)
	rootCmd.Execute()
}

func executeSql(driverName string, connStr string, sqlQuery string) ([]string, error) {
	db, err := sql.Open(driverName, connStr)
	if err != nil {
		return nil, fmt.Errorf("Error opening database:", err)
	}
	defer db.Close()
	return executeFetch(db, sqlQuery)
}

func executeFetch(db *sql.DB, sqlQuery string) ([]string, error) {
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]string, 0)
	for rows.Next() {
		columns, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		datas := make([]string, len(columns))
		var datasPtr []any
		for i := 0; i < len(datas); i++ {
			datasPtr = append(datasPtr, &datas[i])
		}
		if err := rows.Scan(datasPtr...); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		result = append(result, strings.Join(datas, ","))
	}
	return result, nil
}
