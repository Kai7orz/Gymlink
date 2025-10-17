package dto

type ExerciseLikeType struct {
	ExerciseRecordId int64 `json:"exercise_record_id" db:"exercise_record_id"`
}

type RecordCreateType struct {
	ObjectKey      string `json:"object_key" db:"object_key"`
	CleanUpTimeRaw string `json:"clean_up_time" db:"clean_up_time"`
	CleanUpDateRaw string `json:"clean_up_date" db:"clean_up_date"`
	Comment        string `json:"comment" db:"comment"`
}
