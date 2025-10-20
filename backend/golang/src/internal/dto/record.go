package dto

type RecordLikeType struct {
	RecordId int64 `json:"record_id" db:"record_id"`
}

type RecordCreateType struct {
	ObjectKey      string `json:"object_key" db:"object_key"`
	CleanUpTimeRaw string `json:"clean_up_time" db:"clean_up_time"`
	CleanUpDateRaw string `json:"clean_up_date" db:"clean_up_date"`
	Comment        string `json:"comment" db:"comment"`
}
