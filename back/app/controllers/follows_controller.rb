class FollowsController < ApplicationController
  before_action :authenticate_user!

  # POST /api/v1/follows
  def create
    followed = User.find(params[:followed_id]) # followed: フォローされる側

    # 自分自身をフォローできないようにする
    if followed == current_user
      return render json: { error: "Cannot follow yourself" }, status: :unprocessable_entity
    end

    # 既にフォローしているかチェック
    if current_user.following.include?(followed)
      return render json: { error: "Already following this user" }, status: :unprocessable_entity
    end

    follow = current_user.follows_as_follower.build(followed: followed)

    if follow.save
      render json: {message: "Successfully followed user"}, status: :created
    else
      render json: {errors: follow.errors.full_messages}, status: :unprocessable_entity
    end
  end

  # DELETE /api/v1/follows/:followed_id
  def destroy
    followed = User.find(params[:id])
    follow = current_user.follows_as_follower.find_by(followed: followed)

    if follow
      follow.destroy
      render json: { message: "Successfully unfollowed user" }, status: :ok
    else
      render json: { error: "Not following this user" }, status: :not_found
    end
  end
end