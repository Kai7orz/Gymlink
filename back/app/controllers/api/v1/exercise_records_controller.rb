class Api::V1::ExerciseRecordsController < ApplicationController
  before_action :authenticate_user!

  # 運動記録を投稿
  def create
    exercise_record = current_user.exercise_records.build(exercise_record_params)
    if exercise_record.save
      render json: exercise_record, status: :created
    else
      render json: { errors: exercise_record.errors.full_messages }, status: :unprocessable_entity
    end
  end

  # 自分の運動記録一覧を取得
  def my_records
    exercise_records = current_user.exercise_records.order(recorded_date: :desc)
    render json: exercise_records
  end

  # 全ユーザーの運動記録一覧を取得
  def index
    exercise_records = ExerciseRecord.includes(:user).order(recorded_date: :desc)
    records_with_user = exercise_records.map do |record|
      {
        id: record.id,
        exercise_time: record.exercise_time,
        content: record.content,
        recorded_date: record.recorded_date,
        created_at: record.created_at,
        updated_at: record.updated_at,
        user: {
          id: record.user.id,
          name: record.user.name
        }
      }
    end
    render json: records_with_user
  end

  private

  def exercise_record_params
    params.require(:exercise_record).permit(:exercise_time, :content, :recorded_date)
  end
end