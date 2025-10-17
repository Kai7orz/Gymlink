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

type recordLikeTypeDTO struct {
	UId      string `db:"firebase_uid"`
	RecordId int64  `db:"record_id"`
}

func NewExerciseCreateRepo(db *sqlx.DB) *exerciseCreateRepo {
	return &exerciseCreateRepo{db: db}
}

func (r *exerciseCreateRepo) CreateRecordById(ctx context.Context, objectKey string, cleanUpTime int64, cleanUpDate time.Time, comment string, uid string) error {
	type exerciseRawCreateTypeDTO struct {
		UId         string    `db:"firebase_uid"`
		ObjectKey   string    `db:"object_key"`
		CleanUpTime int64     `db:"clean_up_time"`
		CleanUpDate time.Time `db:"clean_up_date"`
		Comment     string    `db:"comment"`
	}
	exerciseCreate := exerciseRawCreateTypeDTO{
		UId:         uid,
		ObjectKey:   objectKey,
		CleanUpTime: cleanUpTime,
		CleanUpDate: cleanUpDate,
		Comment:     comment,
	}
	log.Println("repo create exercise:")
	sql := `INSERT INTO records (user_id,object_key,clean_up_time,clean_up_date,comment) 
	VALUES ((SELECT id FROM users WHERE users.firebase_uid = :firebase_uid LIMIT 1),:object_key,:clean_up_time,:clean_up_date,:comment)
	ON DUPLICATE KEY UPDATE updated_at = NOW();
	`
	_, err := r.db.NamedExec(sql, exerciseCreate)
	if err != nil {
		log.Println("INSERT exercise error: ", err)
		return err
	}
	return nil
}

func (r *exerciseCreateRepo) CreateLike(ctx context.Context, recordId int64, uid string) error {
	exerciseLike := recordLikeTypeDTO{
		UId:      uid,
		RecordId: recordId,
	}
	sql := `INSERT INTO user_likes (user_id,record_id) 
	VALUES ((SELECT id FROM users WHERE users.firebase_uid = :firebase_uid LIMIT 1),:record_id) 
	ON DUPLICATE KEY UPDATE updated_at = NOW()`
	_, err := r.db.NamedExec(sql, exerciseLike)
	if err != nil {
		log.Println("INSERT like error: ", err)
		return err
	}
	return nil
}

func (r *exerciseCreateRepo) DeleteLike(ctx context.Context, recordId int64, uid string) error {
	exerciseDeleteLeike := recordLikeTypeDTO{
		UId:      uid,
		RecordId: recordId,
	}
	sql := `DELETE FROM user_likes WHERE record_id = :record_id and user_id = (SELECT id FROM users WHERE users.firebase_uid = :firebase_uid LIMIT 1)`
	_, err := r.db.NamedExec(sql, exerciseDeleteLeike)
	if err != nil {
		log.Println("DELETE like error: ", err)
		return err
	}
	return nil
}
