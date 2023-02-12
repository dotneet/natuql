package command

import (
	"database/sql"
	"fmt"
	"github.com/dotneet/natuql/index"
	"github.com/dotneet/natuql/openai"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func QueryCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "query",
		Short:   "query to ChatGPT",
		Example: "natuql query \"2022年に一番売れた商品とその合計価格を出力して。\"",
		Run: func(cmd *cobra.Command, args []string) {
			apiKey := viper.GetString("apikey")
			client := openai.NewClient(apiKey)
			query := strings.Join(args, " ")
			schemaIndex, err := index.LoadSchemaIndexFromFile()
			if err != nil {
				// print error to stderr
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				return
			}
			context := schemaIndex.GetRelatedTablesString(query, 8)
			sql, err := client.CreateSql(context, query)
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
