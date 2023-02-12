package openai

import "fmt"

func (client *Client) RefinementSchema(ddl string) (string, error) {
	prompt := fmt.Sprintf(`
DBスキーマにコメントを追加してください。

追加するコメント:
 - テーブル名の日本語訳をコメントして
 - 日本語のテーブル名に別名や類似の名前が考えられる場合は追記して
 - カラム名の日本語訳をコメントして
 - 日本語のカラム名に別名や類似の名前が考えられる場合は追記して

入力例: """
create table drinks(
  id integer,
  -- ドリンク名
  name varchar(255),
  maker varchar(255) COMMENT "メーカー名"
)
"""

出力例： """
-- ドリンクテーブル, 飲料テーブル, 飲み物テーブル
create table drinks(
  -- ドリンクのID, 飲料のID
  id integer,
  -- ドリンク名, ドリンクの名前, 飲料名, 飲料の名前
  name varchar(255),
  -- メーカー名
  maker varchar(255) COMMENT "メーカー名"
)
"""

コメントを追加する対象: """
%s
"""
`, ddl)
	return client.Complete(prompt)
}
