class Api::V1::LikesController < ApplicationController
  before_action :set_record

  # POST /api/v1/exercise_records/:exercise_record_id/likes
  def create
    like = @record.likes.build(user: @current_user)

    if like.save
      render json: { message: "Liked" }, status: :created
    else
      render json: { errors: like.errors.full_messages }, status: :unprocessable_entity
    end
  end

  # DELETE /api/v1/exercise_records/:exercise_record_id/likes
  def delete
    like = @record.likes.find_by(user: @current_user)

    if like
      like.destroy
      render json: { message: "Unliked" }, status: :ok
    else
      render json: { errors: "Not liked found" }, status: :not_found
    end
  end

  private

  def set_record
    @record = ExerciseRecord.find(params[:exercise_record_id])
  end
end