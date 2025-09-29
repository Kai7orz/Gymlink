class AddImageUrlToExerciseRecords < ActiveRecord::Migration[7.2]
  def change
    add_column :exercise_records, :image_url, :string
  end
end
