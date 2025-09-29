class FixFollowsForeignKeys < ActiveRecord::Migration[7.2]
  def change
    # SQLiteの場合、外部キーを変更するにはテーブルを再作成する必要がある
    # 既存のfollowsテーブルをバックアップして再作成する

    # 一時的なテーブル名
    temp_table_name = :follows_temp

    # 既存データを一時テーブルに移動
    execute <<-SQL
      CREATE TABLE #{temp_table_name} AS SELECT * FROM follows;
    SQL

    # 既存のfollowsテーブルを削除
    drop_table :follows

    # 正しい外部キーでfollowsテーブルを再作成
    create_table :follows do |t|
      t.references :follower, null: false, foreign_key: { to_table: :users }
      t.references :followed, null: false, foreign_key: { to_table: :users }
      t.timestamps
    end

    # ユニークインデックスを追加
    add_index :follows, [:follower_id, :followed_id], unique: true

    # データを元に戻す
    execute <<-SQL
      INSERT INTO follows (id, follower_id, followed_id, created_at, updated_at)
      SELECT id, follower_id, followed_id, created_at, updated_at FROM #{temp_table_name};
    SQL

    # 一時テーブルを削除
    drop_table temp_table_name
  end
end
