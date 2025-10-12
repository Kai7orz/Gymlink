package entity

import "time"

type ExerciseRecordType struct {
	Id             int64     `json:"id" db:"id"`
	UserId         int64     `json:"user_id" db:"user_id"`
	UserName       string    `json:"user_name" db:"user_name"`
	Exercise_image string    `json:"exercise_image" db:"exercise_image"`
	ExerciseTime   int64     `json:"exercise_time" db:"exercise_time"`
	ExerciseDate   time.Time `json:"exercise_date" db:"exercise_date"`
	Comment        string    `json:"comment" db:"comment"`
	LikesCount     int64     `json:"likes_count" db:"likes_count"`
}
