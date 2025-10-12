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
