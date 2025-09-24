class AddTeamIdAndPasswordToUsers < ActiveRecord::Migration[7.2]
  def change
    add_reference :users, :team, null: false, foreign_key: true
    add_column :users, :password, :string
  end
end
