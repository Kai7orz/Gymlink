class UserCharactersController < ApplicationController
  before_action :authenticate_user!, only: [:show]

  # GET /user_characters/:user_id
  def show
    user_id = params[:user_id]

    # ユーザーが存在するかチェック
    user = User.find_by(id: user_id)
    return render json: { error: "User not found" }, status: :not_found unless user

    # ユーザーのキャラクター情報を取得
    user_character = UserCharacter.includes(:character).find_by(user: user)
    return render json: { error: "User character not found" }, status: :not_found unless user_character

    # キャラクター情報をレスポンス
    render json: {
      id: user_character.id,
      user_id: user_character.user_id,
      character_id: user_character.character_id,
      level: user_character.level,
      current_points: user_character.current_points,
      required_points: user_character.required_points,
      character_animation: user_character.character_animation
    }
  end
end