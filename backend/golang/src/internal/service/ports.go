package service

import (
	"context"
	"gymlink/internal/entity"

	"firebase.google.com/go/auth"
)

// repository の interface を記述
// 読み取り専用のリポジトリを定義する
type UserQueryRepo interface {
	FindByToken(ctx context.Context, uid string) (*entity.UserType, error)
}

type UserCreateRepo interface {
	CreateUserById(ctx context.Context, name string, avatarUrl string, uid string) error
}

// Firebase とやりとりするためのインターフェース
type AuthClient interface {
	VerifyUser(ctx context.Context, idToken string) (*auth.Token, error)
}

// handler レイヤーが利用するインターフェース
type UserService interface {
	SignUpUser(ctx context.Context, name string, string, idToken string) error
	LoginUser(ctx context.Context, idToken string) (*entity.UserType, error)
}
