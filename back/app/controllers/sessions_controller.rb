class SessionsController < ApplicationController
  # Firebaseトークンでユーザーを認証
  def create
    token = request.headers["Authorization"]&.split(" ")&.last
    payload = FirebaseIdToken::Signature.verify(token)
    return render json: { error: "Invalid token" }, status: :unauthorized unless payload

    firebase_uid = payload["user_id"]

    # DB からユーザー検索
    user = User.find_by(firebase_uid: firebase_uid)

    # レスポンスの分岐
    if user
      render json: { id: user.id, name: user.name }, status: :ok
    else
      render json: { error: "User not found" }, status: :not_found
    end
  end
end
