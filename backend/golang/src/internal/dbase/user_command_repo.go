package dbase

import (
	"context"
	"fmt"
	"gymlink/internal/entity"
	"log"

	"github.com/jmoiron/sqlx"
)

// DB の読み取り群はこのファイルに記述する
type checkFollowTypeDTO struct {
	UId      string `db:"firebase_uid"`
	FollowId int64  `db:"follow_id"`
}

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

func (r *userQueryRepo) CheckFollowById(ctx context.Context, followId int64, uid string) (bool, error) {
	var followed bool

	checkFollow := checkFollowTypeDTO{
		UId:      uid,
		FollowId: followId,
	}
	// users と follows テーブルを連結して，firebase_uid による user_id がfollower_id となっていて，followed_id が followId と合致していればフォローしている
	sql := `
		SELECT EXISTS (
			SELECT 1
			FROM follows
			INNER JOIN users ON users.id = follows.follower_id 
			WHERE follows.followed_id = ?
			AND users.firebase_uid = ?
	)`
	err := r.db.Get(&followed, sql, checkFollow.FollowId, checkFollow.UId)
	if err != nil {
		log.Println("SELECT follow error: ", err)
		return false, err
	}
	return followed, nil
}
