# 認可
# リクエストヘッダーのトークンを常に検証する
class ApplicationController < ActionController::API
  before_action :authenticate_user!

  private

  def authonticate_user!
    token = request.headers["Authorization"]&.spilit(" ")&.last
    payload = FirebaseIdToken::Signature.verify(token)
    return render json: { error: "Unauthorized" }, status: :unauthorized unless payload

    @current_user = User.find_by(firebase_uid: payload["user_id"])
    return render json: { error: "User not found" }, status: :unauthorized unless @current_user
  end
end
