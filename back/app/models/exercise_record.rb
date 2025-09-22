class ExerciseRecord < ApplicationRecord
  belongs_to :user
  has_many :user_likes, dependent: :destroy
end
