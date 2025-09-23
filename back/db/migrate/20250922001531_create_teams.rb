class CreateTeams < ActiveRecord::Migration[7.2]
  def change
    create_table :teams do |t|
      t.references :user, null: false, foreign_key: true
      t.string :name, null: false
      t.date :month, null: false
      t.integer :target_duration, null: false

      t.timestamps
    end
  end
end
