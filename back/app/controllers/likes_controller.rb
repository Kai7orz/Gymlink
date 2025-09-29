class LikesController < ApplicationController
  before_action :authenticate_user!

  # POST /likes
  def create
    exercise_record_id = params[:exercise_record_id]
    exercise_record = ExerciseRecord.find(exercise_record_id)

    # 既存のいいねがあるかチェック
    existing_like = UserLike.find_by(user_id: current_user.id, exercise_record_id: exercise_record.id)
    if existing_like
      return render json: { errors: "Already liked" }, status: :unprocessable_entity
    end

    like = UserLike.new(user: current_user, exercise_record: exercise_record)

    if like.save
      render json: { code: 201, message: "created" }, status: :created
    else
      Rails.logger.error "UserLike save failed: #{like.errors.full_messages}"
      render json: { errors: like.errors.full_messages }, status: :unprocessable_entity
    end
  rescue ActiveRecord::RecordNotFound
    render json: { errors: "Exercise record not found" }, status: :not_found
  end

  # DELETE /likes/:exercise_record_id
  def destroy
    exercise_record = ExerciseRecord.find(params[:exercise_record_id])
    like = exercise_record.user_likes.find_by(user: current_user)

    if like
      like.destroy
      render json: { message: "Unlike successful" }, status: :ok
    else
      render json: { errors: "Like not found" }, status: :not_found
    end
  rescue ActiveRecord::RecordNotFound
    render json: { errors: "Exercise record not found" }, status: :not_found
  end
end