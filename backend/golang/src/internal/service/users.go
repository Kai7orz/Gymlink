package service

import (
	"context"
	"errors"
	"fmt"
	"gymlink/internal/apperrs"
	"gymlink/internal/entity"
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
	if q == nil || cm == nil || p == nil || a == nil {
		return nil, fmt.Errorf("nil error: UserQueryRepo or UserCreateRepo or AuthClient")
	}
	return &userService{q: q, cm: cm, p: p, a: a}, nil
}

func (s *userService) SignUpUser(ctx context.Context, name string, avatarUrl string, idToken string) error {
	//verify user
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return apperrs.ErrVerifyUser
	}

	// token.UID と CharaceterId(デフォルト1) と FirebaseUID と Name と Email を保持
	// v, err := s.q.SignUpById(ctx, 1)
	err = s.cm.CreateUserById(ctx, name, avatarUrl, token.UID)
	if err != nil {
		return fmt.Errorf("failed to create user")
	}

	return nil
}

func (s *userService) LoginUser(ctx context.Context, idToken string) (*entity.UserType, error) {
	//verify user
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return nil, apperrs.ErrVerifyUser
	}

	user, err := s.q.FindByToken(ctx, token.UID)
	if err != nil {
		return nil, fmt.Errorf("error: find user by token : %w", err)
	}
	return user, nil
}

func (s *userService) GetProfile(ctx context.Context, userId int64) (*entity.ProfileType, error) {
	profile, err := s.p.GetProfileById(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get profile by user id : %w", err)
	}
	return profile, nil
}

func (s *userService) FollowUser(ctx context.Context, followerIdRaw int64, followedId int64, idToken string) error {
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return apperrs.ErrVerifyUser
	}
	// firebase Id TOKEN から user_id を取得
	// ユーザーが followerId を偽装していないかなどチェックする
	user, err := s.q.FindByToken(ctx, token.UID)
	if err != nil {
		return fmt.Errorf("error: find by token : %w", err)
	}
	if user.Id != followerIdRaw {
		return fmt.Errorf("invalid: follower must be owner of the account")
	}
	if user.Id == followedId {
		return errors.New("invalid: cannot follow yourself")
	}

	// followerId = user_id とする
	followerId := user.Id

	err = s.cm.FollowUserById(ctx, followerId, followedId)
	if err != nil {
		return fmt.Errorf("failed to follow user by user id : %w", err)
	}

	return nil
}

func (s *userService) GetFollowingById(ctx context.Context, userId int64) ([]entity.UserType, error) {
	following, err := s.q.GetFollowingById(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("cannot get following users : %w", err)
	}
	return following, nil
}

func (s *userService) GetFollowedById(ctx context.Context, userId int64) ([]entity.UserType, error) {
	followed, err := s.q.GetFollowedById(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("cannot get followed users : %w", err)
	}
	return followed, nil
}

func (s *userService) CheckFollowById(ctx context.Context, followId int64, idToken string) (bool, error) {
	//verify user
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return false, apperrs.ErrVerifyUser
	}

	followed, err := s.q.CheckFollowById(ctx, followId, token.UID)
	if err != nil {
		return false, fmt.Errorf("failed to check follow : %w", err)
	}
	return followed, nil
}

func (s *userService) DeleteFollowUser(ctx context.Context, followerIdRaw int64, followedId int64, idToken string) error {
	token, err := s.a.VerifyUser(ctx, idToken)
	if err != nil {
		return apperrs.ErrVerifyUser
	}
	// firebase Id TOKEN から user_id を取得
	// ユーザーが followerId を偽装していないかなどチェックする
	user, err := s.q.FindByToken(ctx, token.UID)
	if err != nil {
		return fmt.Errorf("error: find by token : %w", err)
	}
	if user.Id != followerIdRaw {
		return fmt.Errorf("invalid: follower must be owner of the account")
	}
	if user.Id == followedId {
		return errors.New("invalid: cannot delete yourself")
	}

	// followerId = user_id とする
	followerId := user.Id
	err = s.cm.DeleteFollowUserById(ctx, followerId, followedId)
	if err != nil {
		return fmt.Errorf("failed to delete follow by user id : %w", err)
	}

	return nil

}
