package service

import (
	"gymlink/internal/domain"
	"net/http"
	"time"
)

type fakeToken struct{ UID string }

type fakeAuth struct {
	wantIDToken string
	retUID      string
	err         error
}

type fakeUserQuery struct {
	userByUID      *domain.UserType
	followingUsers []domain.UserType
	followedUsers  []domain.UserType
	findErr        error
	isFollow       bool
	checkErr       error
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

type fakeProfile struct {
	prof *domain.ProfileType
	err  error
}

type fakeRecordQuery struct {
	records  []domain.RecordRawType
	findErr  error
	checkErr error
}

type fakeRecordCmd struct {
	created struct {
		objectKey, cleanTime, likesCount int64
		cleanDate                        time.Time
		comment, uid                     string
		called                           bool
	}
	delete struct {
		recordId, userId int64
		uid              string
		called           bool
	}
	createErr error
	deleteErr error
}

type fakeGptCli struct {
	wantBaseUrl string
	client      *http.Client
}

type fakeAwsCli struct {
	client     string
	bucketName string
}
