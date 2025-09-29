class RenameChracterAnimationInUserCharacters < ActiveRecord::Migration[7.2]
  def change
    rename_column :user_characters, :chracter_animation, :character_animation
  end
end
