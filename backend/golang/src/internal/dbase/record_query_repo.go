package dbase

import (
	"context"
	"fmt"
	"gymlink/internal/entity"

	"github.com/jmoiron/sqlx"
)

type recordQueryRepo struct {
	db *sqlx.DB
}

func NewRecordQueryRepo(db *sqlx.DB) *recordQueryRepo {
	return &recordQueryRepo{db: db}
}

func (r *recordQueryRepo) GetRecordsById(ctx context.Context, id int64) ([]entity.RecordRawType, error) {

	sql := `SELECT 
			records.id,
			records.user_id,
			users.name AS user_name,
			records.object_key,
			records.clean_up_time,
		 	records.clean_up_date,
			records.comment,
			(SELECT COUNT(*) FROM user_likes WHERE record_id = records.id) AS likes_count 
			FROM records INNER JOIN users ON users.id = records.user_id WHERE records.user_id = :id ORDER BY records.id DESC`
	bindParams := map[string]any{
		"id": id,
	}

	query, params, err := sqlx.Named(sql, bindParams)
	if err != nil {
		return nil, fmt.Errorf("sql err : %w", err)
	}
	query = r.db.Rebind(query)
	records := []entity.RecordRawType{}
	if err := r.db.SelectContext(ctx, &records, query, params...); err != nil {
		return nil, fmt.Errorf("select records by user_id=%d: %w", id, err)
	}

	return records, nil

}

func (r *recordQueryRepo) GetRecords(ctx context.Context) ([]entity.RecordRawType, error) {

	sql := `SELECT 
			records.id,
			records.user_id,
			users.name AS user_name,
			records.object_key,
			records.clean_up_time,
			records.clean_up_date,
			records.comment,
			(SELECT COUNT(*) FROM user_likes WHERE record_id = records.id) AS likes_count 
			FROM records INNER JOIN users ON users.id = records.user_id ORDER BY records.id DESC LIMIT 10`
	records := []entity.RecordRawType{}
	if err := r.db.SelectContext(ctx, &records, sql); err != nil {
		return nil, fmt.Errorf("select records : %w", err)
	}
	return records, nil
}
