package dbase

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func SeedingDB(db *sqlx.DB) error {

	type Character struct {
		Id           int64     `json:"id" db:"id"`
		Name         string    `json:"name" db:"name"`
		ImageUrl     string    `json:"image_url" db:"image_url"`
		Level        int64     `json:"level" db:"level"`
		CurrentPoint int64     `json:"current_point" db:"current_point"`
		LimitPoint   int64     `json:"limit_point" db:"limit_point"`
		CreatedAt    time.Time `json:"created_at" db:"created_at"`
		UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	}

	type User struct {
		Id          int64     `json:"id" db:"id"`
		CharacterId int64     `json:"caharacter_id" db:"character_id"`
		FirebaseUid string    `json:"firebase_uid" db:"firebase_uid"`
		Name        string    `json:"name" db:"name"`
		CreatedAt   time.Time `json:"created_at" db:"created_at"`
		UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	}

	type Profile struct {
		Id           int64     `json:"id" db:"id"`
		UserId       int64     `json:"user_id" db:"user_id"`
		ProfileImage string    `json:"profile_image" db:"profile_image"`
		CreatedAt    time.Time `json:"created_at" db:"created_at"`
		UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	}

	type ExerciseRecord struct {
		Id            int64     `json:"id" db:"id"`
		UserId        int64     `json:"user_id" db:"user_id"`
		ExerciseImage string    `json:"exercise_image" db:"exercise_image"`
		ExerciseTime  int64     `json:"exercise_time" db:"exercise_time"`
		ExerciseDate  time.Time `json:"exercise_date" db:"exercise_date"`
		Comment       string    `json:"comment" db:"comment"`
		CreatedAt     time.Time `json:"created_at" db:"created_at"`
		UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	}

	type UserLike struct {
		UserId           int64     `json:"user_id" db:"user_id"`
		ExerciseRecordId int64     `json:"exercise_record_id" db:"exercise_record_id"`
		CreatedAt        time.Time `json:"created_at" db:"created_at"`
		UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	}
	type Follow struct {
		FollowerId int64     `json:"follower_id" db:"follower_id"`
		FollowedId int64     `json:"followed_id" db:"followed_id"`
		CreatedAt  time.Time `json:"created_at" db:"created_at"`
		UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	}

	testCharacter := Character{
		int64(1),
		"Aomaru",
		"character_image.png",
		int64(10),
		int64(5),
		int64(10),
		time.Now(),
		time.Now(),
	}

	testUser := User{
		int64(1),
		int64(1),
		"firebase_test_id",
		"test_user_name",
		time.Now(),
		time.Now(),
	}

	testUser2 := User{
		int64(2),
		int64(1),
		"firebase_test_id2",
		"test_user_name2",
		time.Now(),
		time.Now(),
	}

	testProfile := Profile{
		int64(1),
		int64(1),
		"profile-image.png",
		time.Now(),
		time.Now(),
	}

	testProfile2 := Profile{
		int64(2),
		int64(2),
		"profile-image.png",
		time.Now(),
		time.Now(),
	}

	testExerciseRecord := ExerciseRecord{
		int64(1),
		int64(1),
		"test.png",
		int64(30),
		time.Now(),
		"楽しいけど疲れた!",
		time.Now(),
		time.Now(),
	}

	testExerciseRecord2 := ExerciseRecord{
		int64(2),
		int64(1),
		"test.png",
		int64(10),
		time.Now(),
		"楽しいけど疲れた!",
		time.Now(),
		time.Now(),
	}

	testExerciseRecord3 := ExerciseRecord{
		int64(3),
		int64(2),
		"test.png",
		int64(10),
		time.Now(),
		"楽しい!",
		time.Now(),
		time.Now(),
	}

	testExerciseLike := UserLike{
		int64(1),
		int64(1),
		time.Now(),
		time.Now(),
	}

	testFollow := Follow{
		int64(1),
		int64(2),
		time.Now(),
		time.Now(),
	}

	// 重複データ防止対策は応急処置でしかないので本番環境では他ロジックで対応する
	sql := `INSERT INTO characters (id,name,image_url,level,current_point,limit_point,created_at,updated_at) 
	VALUES (:id,:name,:image_url,:level,:current_point,:limit_point,:created_at,:updated_at) ON DUPLICATE KEY UPDATE name = VALUES(name),updated_at = VALUES(updated_at);`
	bound, err := db.NamedExec(sql, testCharacter)
	log.Println("bound-->", bound)
	if err != nil {
		log.Println("NamedExec Error::", err)
		return err
	}

	sql = `INSERT INTO users (id,character_id,firebase_uid,name,created_at,updated_at) VALUES (:id,:character_id,:firebase_uid,:name,:created_at,:updated_at) ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at);`
	_, err = db.NamedExec(sql, testUser)
	if err != nil {
		log.Println("NamedExec Error::", err)
		return err
	}

	_, err = db.NamedExec(sql, testUser2)
	if err != nil {
		log.Println("NamedExec Error::", err)
		return err
	}

	sql = `INSERT INTO profiles (id,user_id,profile_image,created_at,updated_at) VALUES (:id,:user_id,:profile_image,:created_at,:updated_at) ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at);`
	_, err = db.NamedExec(sql, testProfile)
	if err != nil {
		log.Println("NamedExec Error::", err)
		return err
	}

	sql = `INSERT INTO profiles (id,user_id,profile_image,created_at,updated_at) VALUES (:id,:user_id,:profile_image,:created_at,:updated_at) ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at);`
	_, err = db.NamedExec(sql, testProfile2)
	if err != nil {
		log.Println("NamedExec Error::", err)
		return err
	}

	sql = `INSERT INTO exercise_records (id,user_id,exercise_image,exercise_time,exercise_date,comment,created_at,updated_at) VALUES (:id,:user_id,:exercise_image,:exercise_time,:exercise_date,:comment,:created_at,:updated_at) ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at)`
	_, err = db.NamedExec(sql, testExerciseRecord)
	if err != nil {
		log.Println("NamedExec Error::", err)
		return err
	}

	sql = `INSERT INTO exercise_records (id,user_id,exercise_image,exercise_time,exercise_date,comment,created_at,updated_at) VALUES (:id,:user_id,:exercise_image,:exercise_time,:exercise_date,:comment,:created_at,:updated_at) ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at)`
	_, err = db.NamedExec(sql, testExerciseRecord2)
	if err != nil {
		log.Println("NamedExec Error::", err)
		return err
	}

	sql = `INSERT INTO exercise_records (id,user_id,exercise_image,exercise_time,exercise_date,comment,created_at,updated_at) VALUES (:id,:user_id,:exercise_image,:exercise_time,:exercise_date,:comment,:created_at,:updated_at) ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at)`
	_, err = db.NamedExec(sql, testExerciseRecord3)
	if err != nil {
		log.Println("NamedExec Error::", err)
		return err
	}

	sql = `INSERT INTO user_likes (user_id,exercise_record_id,created_at,updated_at) VALUES (:user_id,:exercise_record_id,:created_at,:updated_at) ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at)`
	_, err = db.NamedExec(sql, testExerciseLike)
	if err != nil {
		log.Println("NamedExec Error::", err)
		return err
	}

	sql = `INSERT INTO follows (follower_id,followed_id,created_at,updated_at) VALUES (:follower_id,:followed_id,:created_at,:updated_at) ON DUPLICATE KEY UPDATE updated_at = VALUES(updated_at)`
	_, err = db.NamedExec(sql, testFollow)
	if err != nil {
		log.Println("NamedExec Error::", err)
		return err
	}

	return nil
}
