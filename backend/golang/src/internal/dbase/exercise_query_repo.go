package dbase

import (
	"context"
	"gymlink/internal/entity"
	"log"

	"github.com/jmoiron/sqlx"
)

type exerciseQueryRepo struct {
	db *sqlx.DB
}

func NewExerciseQueryRepo(db *sqlx.DB) *exerciseQueryRepo {
	return &exerciseQueryRepo{db: db}
}

func (r *exerciseQueryRepo) GetExercisesById(ctx context.Context, id int64) ([]entity.ExerciseRecordType, error) {

	sql := `SELECT 
			exercise_records.id,
			exercise_records.user_id,
			users.name AS user_name,
			exercise_records.exercise_image,
			exercise_records.exercise_time,
			exercise_records.exercise_date,
			exercise_records.comment,
			(SELECT COUNT(*) FROM user_likes WHERE exercise_record_id = exercise_records.id) AS likes_count 
			FROM exercise_records INNER JOIN users ON users.id = exercise_records.user_id WHERE user_id = :id`
	bindParams := map[string]any{
		"id": id,
	}

	query, params, err := sqlx.Named(sql, bindParams)
	if err != nil {
		log.Println("sql err:", err)
	}
	query = r.db.Rebind(query)
	exercises := []entity.ExerciseRecordType{}
	if err := r.db.Select(&exercises, query, params...); err != nil {
		log.Println("err :", err)
		return nil, err
	}

	return exercises, nil

}

func (r *exerciseQueryRepo) GetExercises(ctx context.Context) ([]entity.ExerciseRecordType, error) {

	sql := `SELECT 
			exercise_records.id,
			exercise_records.user_id,
			users.name AS user_name,
			exercise_records.exercise_image,
			exercise_records.exercise_time,
			exercise_records.exercise_date,
			exercise_records.comment,
			(SELECT COUNT(*) FROM user_likes WHERE exercise_record_id = exercise_records.id) AS likes_count 
			FROM exercise_records INNER JOIN users ON users.id = exercise_records.user_id ORDER BY exercise_records.id DESC LIMIT 20`
	exercises := []entity.ExerciseRecordType{}
	if err := r.db.Select(&exercises, sql); err != nil {
		return nil, err
	}
	return exercises, nil
}
