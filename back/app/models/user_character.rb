class UserCharacter < ApplicationRecord
  belongs_to :user
  belongs_to :character

  has_many :character_growths, through: :character
end
