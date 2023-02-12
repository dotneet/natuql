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

```bash
export OPENAI_API_KEY=YOUR_API_KEY
export DATABASE_CONNECTION="root:mysql@tcp(127.0.0.1:3306)/yourdb"
natuql create-index "root:mysql@tcp(127.0.0.1:3306)/yourdb"
natuql query "2022年の売上件数を取得して。"
```
