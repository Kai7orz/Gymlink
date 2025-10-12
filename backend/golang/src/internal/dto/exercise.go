package dto

import "time"

type ExerciseLikeType struct {
	ExerciseRecordId int64 `json:"exercise_record_id" db:"exercise_record_id"`
}

type ExerciseCreateType struct {
	Image        string    `json:"exercise_image" db:"exercise_image"`
	ExerciseTime int64     `json:"exercise_time" db:"exercise_time"`
	Date         time.Time `json:"exercise_date" db:"exercise_date"`
	Comment      string    `json:"comment" db:"comment"`
}
