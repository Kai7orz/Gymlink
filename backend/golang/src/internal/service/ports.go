package service

import (
	"context"
	"gymlink/internal/entity"
	"time"

	"firebase.google.com/go/auth"
)

// repository の interface を記述
// 読み取り専用のリポジトリを定義する
type UserQueryRepo interface {
	FindByToken(ctx context.Context, uid string) (*entity.UserType, error)
}

type UserCreateRepo interface {
	CreateUserById(ctx context.Context, name string, avatarUrl string, uid string) error
	FollowUserById(ctx context.Context, followerId int64, followedId int64) error
}

type ProfileRepo interface {
	GetProfileById(ctx context.Context, id int64) (*entity.ProfileType, error)
}

type ExerciseQueryRepo interface {
	GetExercisesById(ctx context.Context, id int64) ([]entity.ExerciseRecordType, error)
	GetExercises(ctx context.Context) ([]entity.ExerciseRecordType, error)
}

type ExerciseCreateRepo interface {
	CreateExerciseById(ctx context.Context, image string, exerciseTime int64, date time.Time, comment string, uid string) error
	CreateLike(ctx context.Context, exerciseRecordId int64, uid string) error
	DeleteLike(ctx context.Context, exerciseRecordId int64, uid string) error
}

// Firebase とやりとりするためのインターフェース
type AuthClient interface {
	VerifyUser(ctx context.Context, idToken string) (*auth.Token, error)
}

// handler レイヤーが利用するインターフェース
type UserService interface {
	SignUpUser(ctx context.Context, name string, string, idToken string) error
	LoginUser(ctx context.Context, idToken string) (*entity.UserType, error)
	GetProfile(ctx context.Context, id int64) (*entity.ProfileType, error)
	FollowUser(ctx context.Context, followerId int64, followdId int64) error
}

type ExerciseService interface {
	GetExercisesById(ctx context.Context, id int64) ([]entity.ExerciseRecordType, error)
	GetExercises(ctx context.Context) ([]entity.ExerciseRecordType, error)
	CreateExercise(ctx context.Context, image string, exerciseTime int64, date time.Time, comment string, idToken string) error
	CreateLike(ctx context.Context, exerciseRecordId int64, idToken string) error
	DeleteLikeById(ctx context.Context, exerciseRecordId int64, idToken string) error
}
