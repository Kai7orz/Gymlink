package dbase

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type followTypeDTO struct {
	FollowerId int64 `json:"follower_id" db:"follower_id"`
	FollowedId int64 `json:"followed_id" db:"followed_id"`
}

type userCreateTypeDTO struct {
	CharacterId int64  `json:"character_id" db:"character_id"`
	FirebaseUId string `json:"firebase_uid" db:"firebase_uid"`
	Name        string `json:"name" db:"name"`
}

type profileCreateTypeDTO struct {
	FirebaseUId  string `json:"firebase_uid" db:"firebase_uid"`
	Name         string `json:"name" db:"name"`
	ProfileImage string `json:"profile_image" db:"profile_image"`
}

type userCreateRepo struct {
	db *sqlx.DB
}

func NewUserCreateRepo(db *sqlx.DB) *userCreateRepo {
	return &userCreateRepo{db: db}
}

//一通り機能完成後 Create Repo -> Command Repo に変更（作成だけでなく削除の際の機能もこちらに実装するため）

func (r *userCreateRepo) CreateUserById(ctx context.Context, name string, avatarUrl string, uid string) error {

	userCreate := userCreateTypeDTO{
		CharacterId: 1,
		FirebaseUId: uid,
		Name:        name,
	}
	sql := `INSERT INTO users (character_id,firebase_uid,name) VALUES (:character_id,:firebase_uid,:name) ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at);`
	_, err := r.db.NamedExec(sql, userCreate)
	if err != nil {
		log.Println("INSERT user error: ", err)
		return err
	}

	profileCreate := profileCreateTypeDTO{
		FirebaseUId:  uid,
		Name:         name,
		ProfileImage: avatarUrl,
	}

	sql = `INSERT INTO profiles (user_id,profile_image) 
	VALUES ((SELECT id FROM users WHERE users.firebase_uid = :firebase_uid LIMIT 1),:profile_image) ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at);`

	_, err = r.db.NamedExec(sql, profileCreate)
	if err != nil {
		log.Println("INSERT profile error: ", err)
		return err
	}

	return nil
}

func (r *userCreateRepo) FollowUserById(ctx context.Context, follower_id int64, followed_id int64) error {

	userFollow := followTypeDTO{
		FollowerId: follower_id,
		FollowedId: followed_id,
	}
	sql := `INSERT INTO follows (follower_id,followed_id) VALUES (:follower_id,:followed_id) ON DUPLICATE KEY UPDATE updated_at = NOW();`
	_, err := r.db.NamedExec(sql, userFollow)
	if err != nil {
		log.Println("follower_id:", userFollow.FollowerId, " followed_id:", userFollow.FollowedId)
		log.Println("INSERT follow error: ", err)
		return err
	}
	return nil
}

func (r *userCreateRepo) DeleteFollowUserById(ctx context.Context, follower_id int64, followed_id int64) error {
	userDeleteFollow := followTypeDTO{
		FollowerId: follower_id,
		FollowedId: followed_id,
	}
	sql := `DELETE FROM follows WHERE follower_id = :follower_id and followed_id = :followed_id;`
	_, err := r.db.NamedExec(sql, userDeleteFollow)
	if err != nil {
		log.Println("DELETE follow error: ", err)
		return err
	}
	return nil
}
