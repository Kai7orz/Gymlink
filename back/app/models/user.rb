class User < ApplicationRecord
  belongs_to :team

  has_one :profile, dependent: :destroy
  has_many :exercise_records, dependent: :destroy
  has_many :user_likes, dependent: :destroy

  # 自分がフォローしている関係
  has_many :follows_as_follower, class_name: "Follow", foreign_key: "follower_id"
  has_many :following, through: :follows_as_follower, source: :followed


  # 自分をフォローしている関係
  has_many :follows_as_followed, class_name: "Follow", foreign_key: "followed_id"
  has_many :followers, through: :follows_as_followed, source: :follower

end
