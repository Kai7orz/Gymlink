package service

import (
	"context"
	"errors"
	"gymlink/internal/entity"
)

type UserService struct{ q UserQueryRepo }

// UserQueryRepo
func NewUserService(q UserQueryRepo) (*UserService, error) {
	if q == nil {
		return nil, errors.New("nil UserQueryRepo")
	}
	return &UserService{q: q}, nil
}

func (s *UserService) GetUser(ctx context.Context, userId int64) (*entity.UserType, error) {
	if userId <= 0 {
		return nil, errors.New("invalid userId")
	}
	v, err := s.q.FindById(ctx, userId)
	if v == nil {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return v, nil
}
