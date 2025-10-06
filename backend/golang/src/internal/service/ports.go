package service

import (
	"context"
	"gymlink/internal/entity"
)

// repository の interface を記述
// 読み取り専用のリポジトリを定義する
type UserQueryRepo interface {
	FindById(ctx context.Context, id int64) (*entity.UserType, error)
}
