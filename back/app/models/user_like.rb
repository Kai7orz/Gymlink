class UserLike < ApplicationRecord
  belongs_to :user
  belongs_to :exercise_record

  validates :user_id, uniqueness: { scope: :exercise_record_id, message: "has already liked this record" }
end
