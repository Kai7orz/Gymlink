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

func (r *exerciseQueryRepo) GetRecordsById(ctx context.Context, id int64) ([]entity.RecordRawType, error) {

	sql := `SELECT 
			clean_up_records.id,
			clean_up_records.user_id,
			users.name AS user_name,
			clean_up_records.object_key,
			records.clean_up_time,
		 	records.clean_up_date,
			records.comment,
			(SELECT COUNT(*) FROM user_likes WHERE record_id = records.id) AS likes_count 
			FROM records INNER JOIN users ON users.id = records.user_id WHERE user_id = :id`
	bindParams := map[string]any{
		"id": id,
	}

	query, params, err := sqlx.Named(sql, bindParams)
	if err != nil {
		log.Println("sql err:", err)
	}
	query = r.db.Rebind(query)
	records := []entity.RecordRawType{}
	if err := r.db.Select(&records, query, params...); err != nil {
		log.Println("err :", err)
		return nil, err
	}

	return records, nil

}

func (r *exerciseQueryRepo) GetRecords(ctx context.Context) ([]entity.RecordRawType, error) {

	sql := `SELECT 
			records.id,
			records.user_id,
			users.name AS user_name,
			records.object_key,
			records.clean_up_time,
			records.clean_up_date,
			records.comment,
			(SELECT COUNT(*) FROM user_likes WHERE record_id = records.id) AS likes_count 
			FROM records INNER JOIN users ON users.id = records.user_id ORDER BY records.id DESC LIMIT 20`
	records := []entity.RecordRawType{}
	if err := r.db.Select(&records, sql); err != nil {
		return nil, err
	}
	return records, nil
}
