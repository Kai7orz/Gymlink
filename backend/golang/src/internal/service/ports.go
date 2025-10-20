package service

import (
	"context"
	"gymlink/internal/entity"
	"mime/multipart"
	"time"

	"firebase.google.com/go/auth"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// repository の interface を記述
// 読み取り専用のリポジトリを定義する
type UserQueryRepo interface {
	FindByToken(ctx context.Context, uid string) (*entity.UserType, error)
	CheckFollowById(ctx context.Context, followId int64, uid string) (bool, error)
}

type UserCommandRepo interface {
	CreateUserById(ctx context.Context, name string, avatarUrl string, uid string) error
	FollowUserById(ctx context.Context, followerId int64, followedId int64) error
	DeleteFollowUserById(ctx context.Context, followerId int64, followedId int64) error
}

type ProfileRepo interface {
	GetProfileById(ctx context.Context, id int64) (*entity.ProfileType, error)
}

type RecordQueryRepo interface {
	GetRecordsById(ctx context.Context, id int64) ([]entity.RecordRawType, error)
	GetRecords(ctx context.Context) ([]entity.RecordRawType, error)
}

type RecordCommandRepo interface {
	CreateRecordById(ctx context.Context, objectKey string, time int64, date time.Time, comment string, uid string) error
	CreateLike(ctx context.Context, recordId int64, uid string) error
	CheckLike(ctx context.Context, recordId int64, uid string) (bool, error)
	DeleteLike(ctx context.Context, recordId int64, uid string) error
}

// Firebase とやりとりするためのインターフェース
type AuthClient interface {
	VerifyUser(ctx context.Context, idToken string) (*auth.Token, error)
}

type GptClient interface {
	CreateIllustration(ctx context.Context, image *multipart.FileHeader) error
}

type AwsClient interface {
	CheckBucket() error
	UploadImage(ctx context.Context, keyName string) error
	GetObject(
		ctx context.Context,
		objectKey string,
		lifetimeSecs int64) (*v4.PresignedHTTPRequest, error)
}

// handler レイヤーが利用するインターフェース
type UserService interface {
	SignUpUser(ctx context.Context, name string, string, idToken string) error
	LoginUser(ctx context.Context, idToken string) (*entity.UserType, error)
	GetProfile(ctx context.Context, id int64) (*entity.ProfileType, error)
	FollowUser(ctx context.Context, followerId int64, followedId int64) error
	CheckFollowById(ctx context.Context, followId int64, idToken string) (bool, error)
	DeleteFollowUser(ctx context.Context, followerId int64, followedId int64) error
}

type RecordService interface {
	GetRecordsById(ctx context.Context, id int64) ([]entity.RecordType, error)
	GetRecords(ctx context.Context) ([]entity.RecordType, error)
	CreateRecord(ctx context.Context, objectKey string, cleanUpTime string, cleanUpdate string, comment string, idToken string) error
	CreateLike(ctx context.Context, recordId int64, idToken string) error
	CheckLikeById(ctx context.Context, recordId int64, idToken string) (bool, error)
	DeleteLikeById(ctx context.Context, recordId int64, idToken string) error
	GenerateImgAndUpload(ctx context.Context, image *multipart.FileHeader, s3Key string) (string, error)
}
