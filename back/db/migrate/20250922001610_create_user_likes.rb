class CreateUserLikes < ActiveRecord::Migration[7.2]
  def change
    create_table :user_likes do |t|
      t.references :user, null: false, foreign_key: true
      t.references :exercise_record, null: false, foreign_key: true

      t.timestamps
    end

    add_index :user_likes, [:user_id, :exercise_record_id], unique: true
  end
end
