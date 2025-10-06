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
※ 開発環境の場合は -dir の指定策を以下のように dev に変更することに注意
```
migrate create -ext sql -dir ./backend/golang/src/internal/db/migration/dev -seq create_initial_tables　
```


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
モックデータを作成
自動作成の選択肢もあるが今回は直でデータ入れる
参考： https://qiita.com/shion0625/items/e09fe9c008ac6409e57c

DB を開発用と本番用で分ける
1. 新たに開発用の DB を構築
  - 
  ```
    MYSQL_DATABASE=...
    MYSQL_USER=...
    MYSQL_PASSWORD=...
    MYSQL_ROOT_PASSWORD=...

    MYSQL_DEV_DATABASE=gymlink_dev_database
    MYSQL_DEV_USER=gymlink_dev_user
    MYSQL_DEV_PASSWORD=...
  ```
  MySQL の.env に　DEV 用の設定を追加

  コンテナも dev 開発用に建てる．（docker-compose.yml）
```
  dev_db:
    container_name: dev_db
    build:
      context: .
      dockerfile: docker/mysql/Dockerfile
    tty: true
    ports:
      - 3308:3306
    env_file:
      - ./mysql/.env
    environment:
      MYSQL_DATABASE: ${MYSQL_DEV_DATABASE}
      MYSQL_USER: ${MYSQL_DEV_USER}
      MYSQL_DEV_PASSWORD: ${MYSQL_DEV_PASSWORD}
      MYSQL_ROOT_PASSWORD: ${MYSQL_DEV_ROOT_PASSWORD}

    volumes:
      - type: volume
        source: mysql_test_volume
        target: /var/lib/mysql
      - type: bind 
        source: ./mysql/init
        target: /docker-entrypoint-initdb.d
```

2. Go（アプリ側）の DB 接続先を dev 用に変更[ .env に APP_ENV を用意し，その文字列から環境を指定 APP_ENV=development なら dev用DB, APP_ENV=product なら 本番用 DB に接続するロジックに]
3. Go（アプリ側）migration ファイルの指定先の変更[2と同様]
4. INSERT していく 今回は seed ディレクトリに init.go ヲ作成して，アプリ側から初期データを挿入していく方式を採る

※ docker-compose.yml で　開発環境では利用する .env を.dev.env に変更し，本番環境では， .env に

### seed データ挿入にあたって sqlx の導入
```
go get github.com/jmoiron/sqlx
```

*sql.DB を *sqlx.DB に置き換え . また，*sqlx.DB の中に *sql.DB が埋め込まれているので *sqlx.DB の変数に .DB でアクセスして書き換え可能

- データ作成時の createdAt などは time　パッケージの now() などを利用する．

- 依存関係に注意しながらデータ作成（今回はキャラクターが外部依存ないのでキャラクタ―テーブルデータから先に作る必要がある．）


参考：https://github.com/jmoiron/sqlx


db 関連リセットしたいときは docker compose down -v でボリュームごとデータを消去してしまうといい（開発環境においては！！）

### エンドポイントの設置

gin or echo で　お試しでエンドポイントを設置する

今回は学習リソースが豊富な gin を利用する．

エンドポイントの設置にあたって，構成を示しておく．

3層

- transport（handler）
- service（application）
- repository（infra）

handler -> service -> repo interface <- repository
repo interface は service 層に置く．
※ repo interface おいて DB アクセスの内部を知らずにservice はデータ取れるのがメリット
DI によって service の getUser が内部的に変更を加えなくても，service の取得の
する内容（DBからとるのかキャッシュからとるのかモックか）をservice の呼び出し時の
引数（依存性注入部分）の変更だけで切り替えれる．
```
// 本番
svc := service.NewUserService(&dbase.UserDBRepo{db})

// テスト or 開発
svc := service.NewUserService(&MockUserRepo{})

// キャッシュ層を使いたいとき
svc := service.NewUserService(&UserCacheRepo{cache})
```
※はじめはdbaseディレクトリ(=repository)　として実装を進める

### Go で firebase idToken を検証する
参考：https://zenn.dev/minguu42/articles/20220501-go-firebase-auth


参考：

### エンドポイント CRUD の実装

### 重複データ
開発環境では同じデータを何回も流すことがあるので，重複時はINSERT しないといった分岐処理が必要になる．
ON DUPLICATE KEY UPDATE で対処

### swagger の整備
- API 設計

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

#### MySQL Error
2025-10-02T14:30:31.171182Z 1 [ERROR] [MY-012596] [InnoDB] Error number 11 means 'Resource temporarily unavailable'
dev_db          | 2025-10-02T14:30:31.171317Z 1 [ERROR] [MY-012215] [InnoDB] Cannot open datafile './ibdata1'

