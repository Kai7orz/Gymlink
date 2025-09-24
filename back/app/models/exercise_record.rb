class ExerciseRecord < ApplicationRecord
  belongs_to :user
  has_many :user_likes, dependent: :destroy

  validates :exercise_time, presence: true
  validates :content, presence: true
  validates :recorded_date, presence: true
end
