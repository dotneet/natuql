natuql automatically converts natural language inputs to sql and query to a database using the sql.

This tool is proof of concept. Do not use in production.

## Limitation

 - support MySQL only.
 - support japanese language only.

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
natuql query "2022年の売上件数を取得して。"
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

# context-tables-count is the number of tables to be used for query context.
# if set larger value, the query will be more accurate, but it will make the query expensive.
context-tables-count=8
```
