package dbase

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type userCreateRepo struct {
	db *sqlx.DB
}

func NewUserCreateRepo(db *sqlx.DB) *userCreateRepo {
	return &userCreateRepo{db: db}
}

//一通り機能完成後 Create Repo -> Command Repo に変更（作成だけでなく削除の際の機能もこちらに実装するため）

func (r *userCreateRepo) CreateUserById(ctx context.Context, name string, avatarUrl string, uid string) error {
	type userCreateTypeDTO struct {
		CharacterId int64  `json:"character_id" db:"character_id"`
		FirebaseUId string `json:"firebase_uid" db:"firebase_uid"`
		Name        string `json:"name" db:"name"`
	}

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

	return nil
}

func (r *userCreateRepo) FollowUserById(ctx context.Context, follower_id int64, followed_id int64) error {

}
