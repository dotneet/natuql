package index

import (
	"database/sql"
	"fmt"
	"github.com/dotneet/natuql/openai"
	"os"
	"strings"
)

func CreateIndex(client *openai.Client, driverName string, connStr string) error {
	schema, err := loadSchema(driverName, connStr)
	if err != nil {
		return err
	}
	commentedSchema, err := client.RefinementSchema(schema)
	if err != nil {
		return err
	}

	indexFilePath, err := getIndexFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(indexFilePath)
	defer file.Close()
	_, err = file.WriteString(commentedSchema)
	if err != nil {
		return err
	}
	return nil
}

func loadSchema(driverName string, connStr string) (string, error) {
	db, err := sql.Open(driverName, connStr)
	if err != nil {
		return "", fmt.Errorf("Error opening database:", err)
	}
	defer db.Close()
	tables, err := listTables(db)
	if err != nil {
		return "", err
	}
	ddls, err := listDdls(db, tables)
	if err != nil {
		return "", err
	}
	ddl := fmt.Sprintln(strings.Join(ddls, ";\n"))
	return ddl, nil
}

func listTables(db *sql.DB) ([]string, error) {
	// Execute a query
	rows, err := db.Query("show tables")
	if err != nil {
		return nil, fmt.Errorf("error executing query %v", err)
	}
	defer rows.Close()

	// Iterate over the rows
	tables := []string{}
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func listDdls(db *sql.DB, tables []string) ([]string, error) {
	var ddls []string
	for i := 0; i < len(tables); i++ {
		err := (func() error {
			query := fmt.Sprintf("show create table %s", tables[i])
			rows, err := db.Query(query)
			if err != nil {
				return fmt.Errorf("error executing query %v", err)
			}
			defer rows.Close()
			for rows.Next() {
				var table string
				var ddl string
				if err := rows.Scan(&table, &ddl); err != nil {
					return fmt.Errorf("error scanning row: %v", err)
				}
				ddls = append(ddls, ddl)
			}
			return nil
		})()
		if err != nil {
			return nil, err
		}
	}
	return ddls, nil
}
