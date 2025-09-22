class CreateCharacters < ActiveRecord::Migration[7.2]
  def change
    create_table :characters do |t|
      t.string :name, null: false
      t.string :base_image, null: false

      t.timestamps
    end
  end
end
