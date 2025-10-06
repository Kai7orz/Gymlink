package service

import (
	"context"
	"gymlink/internal/entity"

	"firebase.google.com/go/auth"
)

// repository の interface を記述
// 読み取り専用のリポジトリを定義する
type UserQueryRepo interface {
	FindById(ctx context.Context, id int64) (*entity.UserType, error)
}

// Firebase とやりとりするためのインターフェース
type AuthClient interface {
	VerifyUser(ctx context.Context, idToken string) (*auth.Token, error)
}

// handler レイヤーが利用するインターフェース
type UserService interface {
	GetUser(ctx context.Context, idToken string) (*entity.UserType, error)
}
