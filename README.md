natuql enables you to query databases using natural language queries.

This tool is a proof of concept. Do not use in production.

## Features

 - query databases using natural language.
 - supported language: English, Japanese, Chinese, French
 - supported database: MySQL only
 
## Install

```bash
go install github.com/dotneet/natuql@latest
```

## Usage

### Create index and Query

CAUTION: 
OpenAPI is used to build the index.
This is expensive, consuming roughly 2000-4000 tokens per 5 tables.

```bash
export OPENAI_API_KEY=YOUR_API_KEY
export DATABASE_CONNECTION="root:root@tcp(127.0.0.1:3306)/yourdb"
natuql index-create
```

```
natuql query "count users over 30 years old"

Output:
Token usage: 1103.000000
SQL: SELECT COUNT(*) FROM users INNER JOIN profiles ON users.id = profiles.user_id WHERE profiles.age > 30;
Result:
4
```

```
natuql query "2022年の販売件数を取得して。"

Output:
Token usage: 1808.000000
SQL: SELECT COUNT(*) FROM sales WHERE YEAR(sold_at) = 2022;
Result:
2
```

### Rebuilding Index 

If the DB schema is changed, the indexes must be rebuilt.

```bash
natuql index-remove
natuql index-create
```

## Configuration

### Environment Variables

| Name | Description | Default |
| --- | --- |---------|
| OPENAI_API_KEY | OpenAI API Key | empty   |
| DATABASE_CONNECTION | Database connection string | empty   |


### Toml

A toml file must be stored in `$HOME/.config/natuql/config.toml'.

```toml
# Example
apikey=your_secret_key
dbconn=root:root@(tcp:127.0.0.1)/dbname

# language is used for building index file.
# supported language: English, Japanese, Chinese, French
# Default: Auto detection by environment variables.
language=Japanese

# context-tables-count is the number of tables to be used for query context.
# if set larger value, the query will be more accurate, but it will make the query expensive.
# Default: 8
context-tables-count=8
```

