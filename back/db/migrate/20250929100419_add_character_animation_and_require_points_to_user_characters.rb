class AddCharacterAnimationAndRequirePointsToUserCharacters < ActiveRecord::Migration[7.2]
  def change
    add_column :user_characters, :chracter_animation, :string, null: false
    add_column :user_characters, :require_points, :integer, null: false
  end
end
