class Api::V1::UsersController < ApplicationController
  before_action :authenticate_user!, only: [:profile]


  # POST /api/v1/users
  def create
    # Firebase トークンの取得
    token = request.headers["Authorization"]&.split(" ")&.last

    # 認証情報（トークン）が正しいかどうかを確認
    payload = FirebaseIdToken::Signature.verify(token)
    return render json: { error: "Invalid token" }, status: :unauthorized unless payload

    # Firebase UID を取り出す
    firebase_uid = payload["user_id"]

    # UID の重複確認
    if User.exists?(firebase_uid: firebase_uid)
      return render json: { error: "User already exists" }, status: :conflict
    end

    # User インスタンスの作成
    user = User.new(user_params.merge(firebase_uid: firebase_uid))

    # 保存処理
    if user.save
      render json: user, status: :created
    else
      render json: { errors: user.errors.full_messages }, status: :unprocessable_entity
    end
  end


  # GET /api/v1/users/profile
  def profile
    render json: {
      id: current_user.id,
      name: current_user.name,
      profile_image: current_user.avatar_url,
      following_count: current_user.following_count,
      followers_count: current_user.followers_count
    }
  end




  private

  # フロントから送られてくるフォーム入力に対応
  def user_params
    params.require(:user).permit(
      :name,
      :email,
      :password,
    )
  end
end