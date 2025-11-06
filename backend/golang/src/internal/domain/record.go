package domain

import "time"

type RecordType struct {
	Id             int64     `json:"id" db:"id"`
	UserId         int64     `json:"user_id" db:"user_id"`
	UserName       string    `json:"user_name" db:"user_name"`
	PresignedImage string    `json:"presigned_image" db:"presigned_image"`
	CleanUpTime    int64     `json:"clean_up_time" db:"clean_up_time"`
	CleanUpDate    time.Time `json:"clean_up_date" db:"clean_up_date"`
	Comment        string    `json:"comment" db:"comment"`
	LikesCount     int64     `json:"likes_count" db:"likes_count"`
}

type RecordRawType struct {
	Id          int64     `json:"id" db:"id"`
	UserId      int64     `json:"user_id" db:"user_id"`
	UserName    string    `json:"user_name" db:"user_name"`
	ObjectKey   string    `json:"object_key" db:"object_key"`
	CleanUpTime int64     `json:"clean_up_time" db:"clean_up_time"`
	CleanUpDate time.Time `json:"clean_up_date" db:"clean_up_date"`
	Comment     string    `json:"comment" db:"comment"`
	LikesCount  int64     `json:"likes_count" db:"likes_count"`
}
