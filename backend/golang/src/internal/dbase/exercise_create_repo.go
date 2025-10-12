package dbase

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type exerciseCreateRepo struct {
	db *sqlx.DB
}

type exerciseLikeTypeDTO struct {
	UId              string `db:"firebase_uid"`
	ExerciseRecordId int64  `db:"exercise_record_id"`
}

func NewExerciseCreateRepo(db *sqlx.DB) *exerciseCreateRepo {
	return &exerciseCreateRepo{db: db}
}

func (r *exerciseCreateRepo) CreateExerciseById(ctx context.Context, image string, exerciseTime int64, date time.Time, comment string, uid string) error {
	type exerciseRawCreateTypeDTO struct {
		UId          string    `db:"firebase_uid"`
		Image        string    `db:"exercise_image"`
		ExerciseTime int64     `db:"exercise_time"`
		Date         time.Time `db:"exercise_date"`
		Comment      string    `db:"comment"`
	}
	exerciseCreate := exerciseRawCreateTypeDTO{
		UId:          uid,
		Image:        image,
		ExerciseTime: exerciseTime,
		Date:         date,
		Comment:      comment,
	}
	log.Println("repo create exercise:")
	sql := `INSERT INTO exercise_records (user_id,exercise_image,exercise_time,exercise_date,comment) 
	VALUES ((SELECT id FROM users WHERE users.firebase_uid = :firebase_uid LIMIT 1),:exercise_image,:exercise_time,:exercise_date,:comment)
	ON DUPLICATE KEY UPDATE updated_at = NOW();
	`
	_, err := r.db.NamedExec(sql, exerciseCreate)
	if err != nil {
		log.Println("INSERT exercise error: ", err)
		return err
	}
	return nil
}

func (r *exerciseCreateRepo) CreateLike(ctx context.Context, exerciseRecordId int64, uid string) error {
	exerciseLike := exerciseLikeTypeDTO{
		UId:              uid,
		ExerciseRecordId: exerciseRecordId,
	}
	sql := `INSERT INTO user_likes (user_id,exercise_record_id) 
	VALUES ((SELECT id FROM users WHERE users.firebase_uid = :firebase_uid LIMIT 1),:exercise_record_id) 
	ON DUPLICATE KEY UPDATE updated_at = NOW()`
	_, err := r.db.NamedExec(sql, exerciseLike)
	if err != nil {
		log.Println("INSERT like error: ", err)
		return err
	}
	return nil
}

func (r *exerciseCreateRepo) DeleteLike(ctx context.Context, exerciseRecordId int64, uid string) error {
	exerciseDeleteLeike := exerciseLikeTypeDTO{
		UId:              uid,
		ExerciseRecordId: exerciseRecordId,
	}
	sql := `DELETE FROM user_likes WHERE exercise_record_id = :exercise_record_id and user_id = (SELECT id FROM users WHERE users.firebase_uid = :firebase_uid LIMIT 1)`
	_, err := r.db.NamedExec(sql, exerciseDeleteLeike)
	if err != nil {
		log.Println("DELETE like error: ", err)
		return err
	}
	return nil
}
