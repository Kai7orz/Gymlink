## 構成

### Go + MySQL 構築
https://zenn.dev/ajapa/articles/443c396a2c5dd1
を参考に環境を構築.

## 開発の流れ

### golang-migrate の導入
1. ホストマシン側にgolang-migrate を導入（マイグレーション用ファイル作成のため）
```
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

2. go コンテナのためにホストマシン側で以下を実行し go.mod 変更し，go コンテナ起動時に自動で golang-migrate 導入させる

go.mod があるディレクトリに移動してから
```
go get github.com/go-sql-driver/mysql
go get github.com/golang-migrate/migrate/v4
```
でアプリ側でマイグレーションを実行できるようにする.

3. マイグレーションファイルを作成する
- マイグレーションファイルの置き場所となるディレクトリinternal/db/migration の作成（mkdir -p internal/db/migration）

公式ドキュメント
```
migrate create -ext sql -dir db/migrations -seq create_users_table
``` 
を参考に
```
migrate create -ext sql -dir ./backend/golang/src/internal/db/migration -seq create_initial_tables　
```
を実行

マイグレーションファイルが作成されるので up down 両方のファイルを記述
参考： https://qiita.com/tuken24/items/70a3dd8ce0dc3c5a5cce

4. Go アプリ起動に migration 実行させるように実装
```
import (
    "database/sql"
    _ "github.com/lib/pq"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    db, err := sql.Open("postgres", "postgres://localhost:5432/database?sslmode=enable")
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    m, err := migrate.NewWithDatabaseInstance(
        "file:///migrations",
        "postgres", driver)
    m.Up() // or m.Steps(2) if you want to explicitly set the number of migrations to run
}
```
を参考に実装．

- init した場所を取得し，migration ファイルパスの指定を行う（https://zenn.dev/bluetree/articles/1291234ecc6ba1）


golang-migrate公式ドキュメント：
https://github.com/golang-migrate/migrate/tree/master

https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md

5. マイグレーション
マイグレーション時もし，失敗した場合は db が管理している schema_migrations テーブルをいじる必要がある．

db コンテナで MySQL にログインしマイグレートしたいバージョンxを
```
DELETE FROM schema_migrations WHERE version=x;
```
として消して，再度マイグレ

もし中途半端に table できていたら DROP する必要があるのも注意

### テーブルを設計・作成

今回は以下のことを意識

- いつも一緒に読むか・たまにしか読み込まないならテーブル分ける
- 更新時の影響を考えてテーブル分離
- 多対多なら中間テーブル作成
- 役割が異なるかどうか（users と profile テーブル分けるなど）

![alt text](gymlink_table2.png)
profile->profiles (table)
UNIQUE -> PRIMARY (follows)

### Go でテーブルを実装
golang-migrate で 上記設計したテーブルをマイグレーションファイルで実装．
down に関しては依存関係見て DROP 順序気を付ける．
※create TABLE ... は1つのマイグレーションファイルにつき1つにしないとエラーとなる

### Seed データ作成

### Tip・学び
```
services:
  db:
    container_name: db
    image: mysql
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: test
      MYSQL_USER: test
      MYSQL_PASSWORD: test
```
によって， MYSQLにMYSQL_DATABASE で指定した DB へのスーパーユーザー権限が付与される（https://zenn.dev/dotq/articles/54ac397b217f5f）
したがって，開発者がDB 権限をいじらなくても Go -> DB の権限周りで詰まることはない

create TABLE ... は1つのマイグレーションファイルにつき1つにしないとエラーとなる

※ imageキャッシュの意図しないものが利用されるのを防ぐために序盤は docker compose up --build を実行したほうがいい