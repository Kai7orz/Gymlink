class CreateUsers < ActiveRecord::Migration[7.2]
  def change
    create_table :users do |t|
      t.string :firebase_uid, null: false
      t.string :name, null: false
      t.string :email, null: false
      t.string :color, null:false
      t.string :avatar_url

      t.timestamps
    end
  end
end
