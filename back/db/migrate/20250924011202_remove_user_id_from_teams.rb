class RemoveUserIdFromTeams < ActiveRecord::Migration[7.2]
  def change
    remove_reference :teams, :user, null: false, foreign_key: true
  end
end
