package dbase

import (
	"context"
	"gymlink/internal/entity"
	"log"

	"github.com/jmoiron/sqlx"
)

type profileRepo struct {
	db *sqlx.DB
}

func NewProfileRepo(db *sqlx.DB) *profileRepo {
	return &profileRepo{db: db}
}

func (r *profileRepo) GetProfileById(ctx context.Context, id int64) (*entity.ProfileType, error) {

	type followRawTypeDTO struct {
		FollowCountRaw   int64 `json:"follow_count" db:"follow_count"`
		FollowerCountRaw int64 `json:"follower_count" db:"follower_count"`
	}

	type profileRawTypeDTO struct {
		UserId       int64  `json:"user_id" db:"user_id"`
		UserName     string `json:"name" db:"name"`
		ProfileImage string `json:"profile_image" db:"profile_image"`
	}

	sql := "SELECT (SELECT COUNT(*) FROM follows WHERE follower_id = :id) AS follow_count,(SELECT COUNT(*) FROM follows WHERE followed_id= :id ) AS follower_count"
	bindParams := map[string]any{
		"id": id,
	}
	query, params, err := sqlx.Named(sql, bindParams)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var followInfo followRawTypeDTO
	if err := r.db.Get(&followInfo, query, params...); err != nil {
		log.Println("err")
		return nil, err
	}

	sql = "SELECT users.id AS user_id, users.name AS name, profiles.profile_image AS profile_image FROM users INNER JOIN profiles ON users.id = profiles.user_id WHERE users.id = :id"
	// 名前付きパラメタクエリ生成
	query, params, err = sqlx.Named(sql, bindParams)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var profileRaw profileRawTypeDTO
	if err := r.db.Get(&profileRaw, query, params...); err != nil {
		log.Println("err")
		return nil, err
	}

	profile := &entity.ProfileType{
		Id:            profileRaw.UserId,
		Name:          profileRaw.UserName,
		ProfileImage:  profileRaw.ProfileImage,
		FollowCount:   followInfo.FollowCountRaw,
		FollowerCount: followInfo.FollowerCountRaw,
	}
	return profile, nil
}
