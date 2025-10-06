package service

import (
	"context"
	"errors"
	"fmt"
	"gymlink/internal/entity"
)

// service の依存
type userService struct {
	q UserQueryRepo
	a AuthClient
}

// UserQueryRepo
func NewUserService(q UserQueryRepo, a AuthClient) (UserService, error) {
	if q == nil || a == nil {
		return nil, errors.New("nil UserQueryRepo or AuthClient")
	}
	return &userService{q: q, a: a}, nil
}

func (s *userService) GetUser(ctx context.Context, idToken string) (*entity.UserType, error) {
	//verify user
	_, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return nil, errors.New("failed to verify user")
	}

	fmt.Println("Verify User ✅")
	// DB に userId=1 と token 対応したデータをseed しておく
	// 1を idToken からuserId を取り出す処理加えて userId に置き換える
	v, err := s.q.FindById(ctx, 1)
	if v == nil {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return v, nil
}
