class LikesController < ApplicationController
  # POST /likes
  def create
    exercise_record_id = params[:exercise_record_id]
    exercise_record = ExerciseRecord.find(exercise_record_id)
    like = exercise_record.likes.build(user: current_user)

    if like.save
      render json: { code: 201, message: "created" }, status: :created
    else
      render json: { errors: like.errors.full_messages }, status: :unprocessable_entity
    end
  end

  # DELETE /likes/:exercise_record_id
  def destroy
    exercise_record = ExerciseRecord.find(params[:exercise_record_id])
    like = exercise_record.likes.find_by(user: current_user)

    if like
      like.destroy
      render json: { message: "Unlike successful" }, status: :ok
    else
      render json: { errors: "Like not found" }, status: :not_found
    end
  end
end