class CreateExerciseRecords < ActiveRecord::Migration[7.2]
  def change
    create_table :exercise_records do |t|
      t.references :user, null: false, foreign_key: true
      t.integer :exercise_time, null: false
      t.date :recorded_date, null: false
      t.text :content, null: false

      t.timestamps
    end
  end
end
