class ExerciseRecordsController < ApplicationController
  # POST /exercise_records
  def create
    record = ExerciseRecord.new(record_params) # 認証済みユーザーを紐付け
    if record.save
      render json: record, status: :created
    else
      render json: { errors: record.errors.full_messages }, status: :unprocessable_entity
    end
  end

  # GET /exercise_records
  def index
    if params[:user_id]
      records = ExerciseRecord.where(user_id: params[:user_id])
    else
      records = ExerciseRecord.all
    end
    render json: records
  end

  # GET /exercise_records/:id
  def show
    record = ExerciseRecord.find(params[:id])
    render json: record
  rescue ActiveRecord::RecordNotFound
    render json: { error: "Record not found" }, status: :not_found
  end

  # PATCH /exercise_records/:id
  def update
    record = ExerciseRecord.find(params[:id])
    if record.update(record_params)
      render json: record
    else
      render json: { errors: record.errors.full_messages }, status: :unprocessable_entity
    end
  rescue ActiveRecord::RecordNotFound
    render json: { error: "Record not found" }, status: :not_found
  end

  # DELETE /exercise_records/:id
  def destroy
    record = ExerciseRecord.find(params[:id])
    record.destroy!
    render json: { message: "Record deleted successfully", id: record.id }
  rescue ActiveRecord::RecordNotFound
    render json: { error: "Record not found" }, status: :not_found
  end

  private
  def record_params
    params.require(:exercise_record).permit(:exercise_time, :content, :recorded_date, :user_id)
  end
  
  # current_userをfirebaseトークンから取得
  def current_user
  end
end
