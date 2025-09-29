# 認可
# リクエストヘッダーのトークンを常に検証する
class ApplicationController < ActionController::API
  attr_reader :current_user

  private

  def authenticate_user!
    token = request.headers["Authorization"]&.split(" ")&.last
    return render json: { error: "No token provided" }, status: :unauthorized unless token

    begin
      payload = FirebaseIdToken::Signature.verify(token)
      return render json: { error: "Invalid token" }, status: :unauthorized unless payload

      firebase_uid = payload["user_id"]

      # DB からユーザーを検索
      @current_user = User.find_by(firebase_uid: firebase_uid)
      return render json: { error: "User not found" }, status: :unauthorized unless @current_user
    rescue => e
      render json: { error: "Token verification failed" }, status: :unauthorized
    end
  end
end
