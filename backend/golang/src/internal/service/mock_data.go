package service

import (
	"gymlink/internal/entity"
	"time"
)

type fakeToken struct{ UID string }

type fakeAuth struct {
	wantIDToken string
	retUID      string
	err         error
}

type fakeUserQuery struct {
	userByUID *entity.UserType
	findErr   error
	isFollow  bool
	checkErr  error
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

type fakeRecordCmd struct {
	created struct {
		objectKey string
		time      int64
		date      time.Time
		comment   string
		uid       string
		called    bool
	}
	liked struct {
		recordID int64
		uid      string
		called   bool
	}
	unliked struct {
		recordID int64
		uid      string
		called   bool
	}
	deleted struct {
		userID, recordID int64
		uid              string
		called           bool
	}

	createErr error
	deleteErr error
	likeErr   error
	unlikeErr error
	checkLike bool
	checkErr  error
}
