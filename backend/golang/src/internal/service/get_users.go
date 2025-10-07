package service

import (
	"context"
	"errors"
	"gymlink/internal/entity"
	"log"
)

// service の依存
type userService struct {
	q  UserQueryRepo
	cm UserCreateRepo
	a  AuthClient
}

// UserQueryRepo
func NewUserService(q UserQueryRepo, cm UserCreateRepo, a AuthClient) (UserService, error) {
	if q == nil || cm == nil || a == nil {
		return nil, errors.New("nil error: UserQueryRepo or UserCreateRepo or AuthClient")
	}
	return &userService{q: q, cm: cm, a: a}, nil
}

// func (s *userService) GetUser(ctx context.Context, idToken string) (*entity.UserType, error) {
func (s *userService) SignUpUser(ctx context.Context, name string, avatarUrl string, idToken string) error {
	//verify user
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return errors.New("failed to verify user")
	}
	log.Println("UID : ", token.UID)
	log.Println("Verify User ✅")
	// token.UID と CharaceterId(デフォルト1) と FirebaseUID と Name と Email を保持

	// v, err := s.q.SignUpById(ctx, 1)
	err = s.cm.CreateUserById(ctx, name, avatarUrl, token.UID)
	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (s *userService) LoginUser(ctx context.Context, idToken string) (*entity.UserType, error) {
	//verify user
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return nil, errors.New("failed to verify user")
	}
	log.Println("Verify User ✅")

	user, err := s.q.FindByToken(ctx, token.UID)
	if err != nil {
		log.Println("failed to find user by token")
		return nil, err
	}
	return user, nil
}
