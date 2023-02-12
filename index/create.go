package index

import (
	"database/sql"
	"fmt"
	"github.com/dotneet/natuql/openai"
	"github.com/dotneet/natuql/path"
	"strings"
)

func CreateIndex(client *openai.Client, driverName string, connStr string) error {
	schema, err := loadSchema(driverName, connStr)
	if err != nil {
		return err
	}

	chunks := schema.Chunks(5)
	var commentedSchemas []string
	separator := "/*SPECIAL_SEPARATOR*/"
	for i := 0; i < len(chunks); i++ {
		chunk := chunks[i]
		commentedSchema, err := client.RefineSchema(strings.Join(chunk, ";\n"+separator+"\n"))
		if err != nil {
			return err
		}
		schemas := strings.Split(commentedSchema, separator)
		commentedSchemas = append(commentedSchemas, schemas...)
	}

	schemaIndex := &SchemaIndex{
		Tables: commentedSchemas,
	}

	indexFilePath, err := path.GetIndexFilePath()
	if err != nil {
		return err
	}
	schemaIndex.WriteAsJson(indexFilePath)
	return nil
}

func loadSchema(driverName string, connStr string) (*Schema, error) {
	db, err := sql.Open(driverName, connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database. %v", err)
	}
	defer db.Close()
	tables, err := listTables(db)
	if err != nil {
		return nil, err
	}
	ddls, err := listDdls(db, tables)
	if err != nil {
		return nil, err
	}
	return &Schema{
		tables: ddls,
	}, nil
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
