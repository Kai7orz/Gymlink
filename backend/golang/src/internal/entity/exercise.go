package entity

import "time"

type ExerciseRecordType struct {
	Id             int64     `db:"id"`
	UserId         int64     `db:"user_id"`
	UserName       string    `db:"user_name"`
	Exercise_image string    `db:"exercise_image"`
	ExerciseTime   int64     `db:"exercise_time"`
	ExerciseDate   time.Time `db:"exercise_date"`
	Comment        string    `db:"comment"`
	LikesCount     int64     `db:"likes_count"`
}
