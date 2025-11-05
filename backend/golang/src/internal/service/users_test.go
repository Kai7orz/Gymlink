package service

import (
	"context"
	"gymlink/internal/domain"
	"gymlink/internal/mocks"
	"testing"

	"firebase.google.com/go/auth"
)

// テストコードはAI にベースを書かせて, それを読んで（写経）理解・活用の流れをとってます

type fakeToken struct{ UID string }

type fakeAuth struct {
	wantIDToken string
	retUID      string
	err         error
}

func (f *fakeAuth) VerifyUser(ctx context.Context, idToken string) (*auth.Token, error) {
	if f.err != nil {
		return nil, f.err
	}
	return &auth.Token{UID: f.retUID}, nil
}

type fakeUserQuery struct {
	userByUID *domain.UserType
	findErr   error
	isFollow  bool
	checkErr  error
}

func (f *fakeUserQuery) FindByToken(ctx context.Context, uid string) (*domain.UserType, error) {
	if f.findErr != nil {
		return nil, f.findErr
	}
	return f.userByUID, nil
}

func (f *fakeUserQuery) CheckFollowById(ctx context.Context, followId int64, uid string) (bool, error) {
	if f.findErr != nil {
		return false, f.checkErr
	}
	return f.isFollow, nil
}

type fakeUserCmd struct {
	created struct {
		name, avatar, uid string
		called            bool
	}
	follow struct {
		follower, followed int64
		called             bool
	}
	unfollow struct {
		follower, followed int64
		called             bool
	}
	createErr   error
	followErr   error
	unfollowErr error
}

func (f *fakeUserCmd) CreateUserById(ctx context.Context, name, avatarURL, uid string) error {
	if f.createErr != nil {
		return f.createErr
	}
	f.created = struct {
		name   string
		avatar string
		uid    string
		called bool
	}{name, avatarURL, uid, true}
	return nil
}

func (f *fakeUserCmd) FollowUserById(ctx context.Context, followerId, followedId int64) error {
	if f.followErr != nil {
		return f.followErr
	}
	f.follow = struct {
		follower int64
		followed int64
		called   bool
	}{
		followerId,
		followedId,
		true,
	}
	return nil
}

func (f *fakeUserCmd) DeleteFollowUserById(ctx context.Context, followerId, followedId int64) error {
	if f.unfollowErr != nil {
		return f.unfollowErr
	}
	f.unfollow = struct {
		follower int64
		followed int64
		called   bool
	}{followerId, followedId, true}
	return nil
}

type fakeProfile struct {
	prof *domain.ProfileType
	err  error
}

func (f *fakeProfile) GetProfileById(ctx context.Context, id int64) (*domain.ProfileType, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.prof, nil
}

func TestNewUserService_NilDeps(t *testing.T) {
	_, err := NewUserService(nil, &fakeUserCmd{}, &fakeProfile{}, &fakeAuth{})
	if err == nil {
		t.Fatal("want error where q is nil")
	}

	_, err = NewUserService(&mocks.FakeUserQuery{}, nil, &fakeProfile{}, &fakeAuth{})
	if err == nil {
		t.Fatal("want error where cm is nil")
	}

	_, err = NewUserService(&mocks.FakeUserQuery{}, &fakeUserCmd{}, nil, &fakeAuth{})
	if err == nil {
		t.Fatal("want error where p is nil")
	}

	_, err = NewUserService(&mocks.FakeUserQuery{}, &fakeUserCmd{}, &fakeProfile{}, nil)
	if err == nil {
		t.Fatal("want error where q is nil")
	}
}

func TestSignUpUser_Success(t *testing.T) {
	q := &fakeUserQuery{}
	cm := &fakeUserCmd{}
	p := &fakeProfile{}
	a := &fakeAuth{retUID: "u-123"}

	svc, _ := NewUserService(q, cm, p, a)
	err := svc.SignUpUser(context.Background(), "TestUser", "avatar_url", "idToken")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	if !cm.created.called {
		t.Fatalf("CreateUserById not called")
	}
	if cm.created.uid != "u-123" || cm.created.name != "TestUser" {
		t.Fatalf("args not passed correctly: %+v", cm.created)
	}
}
