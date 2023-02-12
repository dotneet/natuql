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
