class UsersController < ApplicationController


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


  # GET /user_profiles/:user_id
  def profile
    user = User.find_by(id: params[:user_id])
    return render json: { error: "User not found" }, status: :not_found unless user

    render json: {
      id: user.id,
      name: user.name,
      profile_image: user.avatar_url,
      following_count: user.following_count,
      followers_count: user.followers_count
    }
  end




  private

  # フロントから送られてくるフォーム入力に対応
  def user_params
    params.require(:user).permit(
      :name,
      :avatar_url,
    )
  end
end