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
そしてその後は docker compose up --build でキャッシュ無視してビルドする

### エンドポイントの設置

gin or echo で　お試しでエンドポイントを設置する

今回は学習リソースが豊富な gin を利用する．

エンドポイントの設置にあたって，構成を示しておく．

- transport（handler）
- service（application）
- repository（infra）
- adapter（外部依存）

handler -> service -> repo interface <- repository,adapter

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
※service が外部サービスとDB　をインターフェースで扱うように定める
### Go で firebase idToken を検証する

- service に auth（外部サービス）用のインターフェス定義
- main.go で注入
- auth の具体を実装

参考：https://zenn.dev/minguu42/articles/20220501-go-firebase-auth

### エンドポイント CRUD の実装

JOIN 使ってデータ取得：

### 重複データ
開発環境では同じデータを何回も流すことがあるので，重複時はINSERT しないといった分岐処理が必要になる．
ON DUPLICATE KEY UPDATE で対処

### DB からのデータ取得
https://qiita.com/k-motoyan/items/f37d1348efd3f40e9096

API 実装の手順
例）
exercise 実装流れ
1. entity Exercise を定義
2. handler に Exercise 取得のためのハンドラーを実装
3. service に repo interface を利用して，Exercise 情報を取得
4. repo interface で Exercise 情報取得する関数を定義
5. repository で実際の取得処理を実装

### テーブル結合の基礎・SQL
フォロー・フォロワー数取得 SQL など
https://hit.hateblo.jp/entry/2016/05/09/131806

参考：https://note.shiftinc.jp/n/nd2a2915211f8


### curl で firebase の idToken 取得して，authorization つけた状態でリクエスト
参考：https://zenn.dev/purratto/articles/94693409a8c62b
```
curl -X POST -H "Content-Type: application/json" -d '{
  "email": "<EMAIL>",
  "password": "<PASSWORD>",
  "returnSecureToken": true
}' "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=AIzaSyDLENAgDartCqlupd9YZTVAylQ6XQ8Ivw4"
```
で idToken を取得し
```
 curl -H GET 'http://localhost:8080/user_profiles/1' -H 'Content-Type:application/json;charset=utf-8' -H 'Authorization: Bearer 取得したidToken
```
でリクエストを送信．
### テーブル変更（DB リセットできる開発環境時の方法）
- exercise のカラムに exercise_image を追加
  - exercise テーブルをCreate する table create 用のマイグレーションファイルを変更
  - データをINSERT するマイグレーションファイルの変更
  - exercise テーブルのデータを取得する repository の変更
  - exercise テーブルの struct 変更
  
### like 機能の実装
- like 機能は現状exercise の記録にしかつけない予定で、コメントへのいいね機能など拡張の予定ないため，ロジックを like 専用に切り出しはしない．exercise ファイルに同居させる（exercise のCreate 側のロジックに組み込む）.

### follow 機能の実装
- like 機能のロジックを転用
  - handler に /follows を設置
  - handler で request Body から follower_id , followed_id を読み込む
  - service で follower_id , followed_id をインターフェース通じて Repo に渡す
  - Repo で follows　テーブルに INSERT

### DTO 
- handler で必要な DTO は dto ディレクトリにまとめる
- repository では，各関数または同一ファイル内に DTO を直接宣言しておく．（database に直接触れる部分ではチェックを厳しくしたい　＆　構造をすぐに把握できるようにしたい　ただコード肥大化するので将来的には dto ディレクトリに全部まとめるかも）
### swagger の整備
- API 設計

### エラー集
- follow 機能で存在しないユーザーへのフォロー
```
go              | 2025/10/12 08:15:39 follower_id: 4  followed_id: 0
go              | 2025/10/12 08:15:39 INSERT follow error:  Error 1452 (23000): Cannot add or update a child row: a foreign key constraint fails (`gymlink_dev_database`.`follows`, CONSTRAINT `follows_ibfk_2` FOREIGN KEY (`followed_id`) REFERENCES `users` (`id`) ON DELETE CASCADE)
go              | 2025/10/12 08:15:39 failed to follow user by user id  Error 1452 (23000): Cannot add or update a child row: a foreign key constraint fails (`gymlink_dev_database`.`follows`, CONSTRAINT `follows_ibfk_2` FOREIGN KEY (`followed_id`) REFERENCES `users` (`id`) ON DELETE CASCADE)
go              | 2025/10/12 08:15:39 error: user follow
go              | [GIN] 2025/10/12 - 08:15:39 | 400 |    3.9
```
```
mysql> select * from users;
+----+--------------+------------------------------+-----------------+---------------------+---------------------+
| id | character_id | firebase_uid                 | name            | created_at          | updated_at          |
+----+--------------+------------------------------+-----------------+---------------------+---------------------+
|  1 |            1 | firebase_test_id             | test_user_name  | 2025-10-09 13:16:21 | 2025-10-12 08:14:09 |
|  2 |            1 | firebase_test_id2            | test_user_name2 | 2025-10-09 13:16:21 | 2025-10-12 08:14:09 |
|  4 |            1 | BnarbqOmmoOHbCGcn09ernBalCt2 | kai7orz         | 2025-10-11 21:26:02 | 2025-10-11 21:26:02 |
+----+--------------+------------------------------+-----------------+---------------------+---------------------+
3 rows in set (0.00 sec)

mysql>
```

- まずは DB 内に id がいくつのユーザーがいるかをチェック
- DB user 内にいないのに exercise の user_id で使われてしまっているid がいれば user を seeding のときに追加する
- followed_id = 0　となる原因の追究
- user_id = 2の user はinsert しているが，profile テーブルに user_id = 2 profile を設定していないため，カードリストからプロフィール画面に飛べず，フォローできない問題がある．
  - seeding すれば解決（実運用の際には問題発生しなさそう）

```
mysql> select * from exercise_records ;
+----+---------+------------------------+---------------+---------------------+---------------------------+---------------------+---------------------+
| id | user_id | exercise_image         | exercise_time | exercise_date       | comment                   | created_at          | updated_at          |
+----+---------+------------------------+---------------+---------------------+---------------------------+---------------------+---------------------+
|  1 |       1 | test.png               |            30 | 2025-10-09 13:16:21 | 楽しいけど疲れた!         | 2025-10-09 13:16:21 | 2025-10-12 08:14:09 |
|  2 |       1 | test.png               |            10 | 2025-10-09 13:22:07 | 楽しいけど疲れた!         | 2025-10-09 13:22:07 | 2025-10-12 08:14:09 |
|  3 |       2 | test.png               |            10 | 2025-10-09 14:02:58 | 楽しい!                   | 2025-10-09 14:02:58 | 2025-10-12 08:14:09 |
|  4 |       4 |                        |             0 | 2025-09-26 15:00:00 |                           | 2025-10-11 22:14:09 | 2025-10-11 22:14:09 |
|  5 |       4 | /images/sportImage.png |            30 | 2025-09-26 15:00:00 | nice                      | 2025-10-12 04:33:22 | 2025-10-12 04:33:22 |
+----+---------+------------------------+---------------+---------------------+---------------------------+---------------------+---------------------+
5 rows in set (0.00 sec)

mysql>
```

followed_id = 0 となる問題は，追及する必要がありそう（DB 内に user_id = 0 のレコードがないのにもかかわらず，フロントから followed_id = 0 のデータが流れてきているのはバグ）
  - フロント側でどのように user_id を指定して post してきているかをチェックする．
  - follow post 時のuser_id は cardlist の user_id を参照していたはずなので， follow + cardlist あたりのコードを調査していく．
    - Profile.vue から流れてきた id を followd_id としてpost している．
    - Profile の id は
    ```
        const data = await $fetch(
        `/api/user_profiles/${uid}`,{
          headers: {
                        'Authorization': 'Bearer ' + TOKEN,
                        'Content-Type': 'application/json'
                    },
        }
    );

    ```
    の data からとってきていて，
    ```
        const follow = async (id:number) => {
        await navigateTo('/following')
        console.log("follower_id:",user.userId)
        console.log("followed_id:",data.id)
        await $fetch("/api/follows", {
                        method: 'POST',
                        headers: {
                            'Authorization': 'Bearer ' + TOKEN,
                            'Content-Type': 'application/json'
                        },
                        body: {
                            follower_id: user.userId,
                            followed_id: id
                        },
                    })
    }
  ```
  として follow に post しているが， 
  解決： backend では json タグ付けていなくて data　の id を Id として返していたが， フロント側は id を期待してコード書いていた．
  backend に json タグ付けて対応


  - 応急処置として seeding.go で profile 増やす．



- FireBase のエラー
  - ユーザー・パスワード登録直後はユーザー認証通るが，一定時間経過後にログインしようとすると invalid credential が発生する．リフレッシュの影響？？
  ```
  go              | 2025/10/11 21:53:23 error:  failed to verify token signature
  ```
  - curl で試していたが，アプリ側に実装して試す
  これでもし アプリ側に問題がなければ，curl で取得する idToken が間違っている可能性あり
    - アプリ側で post したら問題なく通ったが， curl で上記手順でidToken 取得して POST したら failed to verify user error なので curl で取得してきた idToken の渡し方間違っている可能性ある．本来のユースケースであるアプリ側からの動作は問題ないのでこのまま開発をすすめて，完成後に修正する（issue 化済み）

- /user_profiles/:user_id で 多分 follower_id に登録されていないid を user_id に指定するとエラーになる．本来 follower_id に登録されていなくても
followed_id に登録されていたら正常にレスポンスしたいのにできていない

- main の DI で
```
	exerciseSvc, err := service.NewExerciseService(exerciseQueryRepo, exerciseCreateRepo, authC)
	if err != nil {
		log.Fatal("exercise service error")
	}
```
をしているのに， authC が nil error で詰まった．
原因は service のコンストラクタで func NewExerciseService(e ExerciseQueryRepo, ec ExerciseCreateRepo, a AuthClient) (ExerciseService, error) までは定義していたが，
return &exerciseService{e: e}, nil で実態を exerciseCreateRepoo, authC 返すの忘れていたためエラーが出ていた．

```
package service

// エラー原因となったコード

import (
	"context"
	"errors"
	"gymlink/internal/entity"
	"log"
	"time"
)

type exerciseService struct {
	e  ExerciseQueryRepo
	ec ExerciseCreateRepo
	a  AuthClient
}

func NewExerciseService(e ExerciseQueryRepo, ec ExerciseCreateRepo, a AuthClient) (ExerciseService, error) {
	if e == nil {
		return nil, errors.New("nil error: ExerciseQueryRepo or ExerciseCreateRepo")
	}
	return &exerciseService{e: e}, nil
}

func (s *exerciseService) GetExercisesById(ctx context.Context, id int64) ([]entity.ExerciseRecordType, error) {
	exercises, err := s.e.GetExercisesById(ctx, id)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return exercises, nil
}

func (s *exerciseService) GetExercises(ctx context.Context) ([]entity.ExerciseRecordType, error) {
	exercises, err := s.e.GetExercises(ctx)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	return exercises, nil
}

func (s *exerciseService) CreateExercise(ctx context.Context, image string, exerciseTime int64, date time.Time, comment string, idToken string) error {
	if s.a == nil {
		log.Println("auth :nil ✅")
		return errors.New("eror auth nil")
	}
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return errors.New("failed to verify user")
	}
	err = s.ec.CreateExerciseById(ctx, image, exerciseTime, date, comment, token.UID)
	if err != nil {
		log.Println("error: ", err)
		return err
	}
	return nil
}

```

authorization が verify してくれない　泣

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

