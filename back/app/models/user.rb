class User < ApplicationRecord
  belongs_to :team

  has_one :profile, dependent: :destroy
  has_one :user_character, dependent: :destroy
  has_many :exercise_records, dependent: :destroy
  has_many :user_likes, dependent: :destroy

  # 自分がフォローしている関係
  has_many :follows_as_follower, class_name: "Follow", foreign_key: "follower_id", dependent: :destroy
  has_many :following, through: :follows_as_follower, source: :followed


  # 自分をフォローしている関係
  has_many :follows_as_followed, class_name: "Follow", foreign_key: "followed_id", dependent: :destroy
  has_many :followers, through: :follows_as_followed, source: :follower

  validates :name, presence: true

  # フォロー数を取得
  def following_count
    following.count
  end

  # フォロワー数を取得
  def followers_count
    followers.count
  end

end
