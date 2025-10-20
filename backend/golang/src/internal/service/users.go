package service

import (
	"context"
	"errors"
	"gymlink/internal/entity"
	"log"
)

// service の依存
type userService struct {
	q  UserQueryRepo   // ユーザー読み取り用 repo interface
	cm UserCommandRepo // ユーザー書き込み用 repo interface
	p  ProfileRepo
	a  AuthClient // 外部Auth との接続用 interface
}

// UserQueryRepo
func NewUserService(q UserQueryRepo, cm UserCommandRepo, p ProfileRepo, a AuthClient) (UserService, error) {
	if q == nil || cm == nil || a == nil {
		return nil, errors.New("nil error: UserQueryRepo or UserCreateRepo or AuthClient")
	}
	return &userService{q: q, cm: cm, p: p, a: a}, nil
}

func (s *userService) SignUpUser(ctx context.Context, name string, avatarUrl string, idToken string) error {
	//verify user
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return errors.New("failed to verify user")
	}

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

	user, err := s.q.FindByToken(ctx, token.UID)
	if err != nil {
		log.Println("failed to find user by token")
		return nil, err
	}
	return user, nil
}

func (s *userService) GetProfile(ctx context.Context, id int64) (*entity.ProfileType, error) {
	profile, err := s.p.GetProfileById(ctx, id)
	if err != nil {
		log.Println("failed to get profile by user id ", err)
		return nil, err
	}
	return profile, nil
}

func (s *userService) FollowUser(ctx context.Context, followerId int64, followedId int64) error {
	err := s.cm.FollowUserById(ctx, followerId, followedId)
	if err != nil {
		log.Println("failed to follow user by user id ", err)
		return err
	}

	return nil
}

func (s *userService) CheckFollowById(ctx context.Context, followId int64, idToken string) (bool, error) {
	//verify user
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return false, errors.New("failed to verify user")
	}

	followed, err := s.q.CheckFollowById(ctx, followId, token.UID)
	if err != nil {
		log.Println("error: check follow ", err)
		return false, err
	}
	return followed, nil
}

func (s *userService) DeleteFollowUser(ctx context.Context, followerId int64, followedId int64) error {
	err := s.cm.DeleteFollowUserById(ctx, followerId, followedId)
	if err != nil {
		log.Println("failed to delete follow by user id", err)
		return err
	}

	return nil

}
