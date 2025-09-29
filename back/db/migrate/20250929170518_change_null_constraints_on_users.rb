class ChangeNullConstraintsOnUsers < ActiveRecord::Migration[7.2]
  def change
    change_column_null :users, :firebase_uid, true
    change_column_null :users, :color, true
    change_column_null :users, :team_id, true
  end
end
