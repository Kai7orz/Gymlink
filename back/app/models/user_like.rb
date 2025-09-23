class UserLike < ApplicationRecord
  belongs_to :user
  belongs_to :exercise_record
end
