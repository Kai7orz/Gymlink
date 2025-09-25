class Api::v1::UsersController < ApplicationController
  # POST /api/v1/users (with Firebase token)
  def create
    token = request.headers["Authorization"]&.split(" ")&.last
    payload = FirebaseIdToken::Signature.verify(token)
    return render json: { error: "Invalid token" }, status: :unauthorized unless payload

    firebase_uid = payload["user_id"]

    # 新規登録専用: UID が既にあればエラー
    if User.exists?(firebase_uid: firebase_uid)
      return render json: { error: "User already exists" }, status: :conflict
    end
    user = User.new(user_params.merge(firebase_uid: firebase_uid))
    if user.save
      render json: user, status: :created
    else
      render json: { errors: user.errors.full_messages }, status: :unprocessable_entity
    end
  end

  private

  def user_params
    params.require(:user).permit(
      :name
      :email
      :avatar_url
      :password
      :team_id
      :color
    )
  end
end