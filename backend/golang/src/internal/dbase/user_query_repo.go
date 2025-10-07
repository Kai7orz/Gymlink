package dbase

import (
	"context"
	"fmt"
	"gymlink/internal/entity"

	"github.com/jmoiron/sqlx"
)

// DB の読み取り群はこのファイルに記述する

type userQueryRepo struct {
	db *sqlx.DB
}

func NewUserQueryRepo(db *sqlx.DB) *userQueryRepo {
	return &userQueryRepo{db: db}
}

func (r *userQueryRepo) FindByToken(ctx context.Context, uid string) (*entity.UserType, error) {
	sql := `SELECT id,name FROM users WHERE firebase_uid=?`
	var user entity.UserType
	if err := r.db.GetContext(ctx, &user, sql, uid); err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}
	return &user, nil
}
