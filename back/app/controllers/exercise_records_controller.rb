class ExerciseRecordsController < ApplicationController
  before_action :authenticate_user!

  # 運動記録を投稿
  def create
    exercise_record = current_user.exercise_records.build(exercise_record_params)
    if exercise_record.save
      record_with_details = {
        id: exercise_record.id,
        user_id: exercise_record.user.id,
        user_name: exercise_record.user.name,
        image_url: "/images/sportImage.png",
        time: exercise_record.exercise_time,
        date: exercise_record.recorded_date,
        comment: exercise_record.content,
        likes_count: 0
      }
      render json: record_with_details, status: :created
    else
      render json: { errors: exercise_record.errors.full_messages }, status: :unprocessable_entity
    end
  end

  # 自分の運動記録一覧を取得
  def user_exercises
    exercise_records = current_user.exercise_records.includes(:user, :user_likes).order(recorded_date: :desc)
    records_with_details = exercise_records.map do |record|
      {
        id: record.id,
        user_id: record.user.id,
        user_name: record.user.name,
        image_url: "/images/sportImage.png",
        time: record.exercise_time,
        date: record.recorded_date,
        comment: record.content,
        likes_count: record.user_likes.count
      }
    end
    render json: records_with_details
  end

  # 全ユーザーの運動記録一覧を取得（最新20件）
  def index
    exercise_records = ExerciseRecord.includes(:user, :user_likes).order(recorded_date: :desc).limit(20)
    records_with_details = exercise_records.map do |record|
      {
        id: record.id,
        user_id: record.user.id,
        user_name: record.user.name,
        image_url: "/images/sportImage.png",
        time: record.exercise_time,
        date: record.recorded_date,
        comment: record.content,
        likes_count: record.user_likes.count
      }
    end
    render json: records_with_details
  end

  private

  def exercise_record_params
    params.require(:exercise_record).permit(:exercise_time, :content, :recorded_date)
  end
end