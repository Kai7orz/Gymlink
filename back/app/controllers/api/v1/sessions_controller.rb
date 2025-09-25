# Firebaseトークンでユーザーを認証し、アプリ内のセッション or JWTを発行
# POST /api/v1/sessions
def create
  token = request.headers["Authorization"]&.split(" ")&.last
  payload = FirebaseIdToken::Signature.verify(token)
  return render json: { error: "Invalid token" }, status: :unauthorized unless payload

  firebase_uid = payload["user_id"]
  user = User.find_by(firebase_uid: firebase_uid)

  if user
    render json: { message: "Login successful", user: user }, status: :omakase
  else
    render json: { error: "User not found" }, status: :not_found
  end
end