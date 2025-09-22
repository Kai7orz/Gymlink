class CreateCharacterGrowths < ActiveRecord::Migration[7.2]
  def change
    create_table :character_growths do |t|
      t.references :character, null: false, foreign_key: true
      t.integer :level, null: false
      t.integer :required_points, null: false
      t.string :reward_image_url, null: false

      t.timestamps
    end

    add_index :character_growths, [:character_id, :level], unique: true
  end
end
