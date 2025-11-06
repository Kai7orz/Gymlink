package domain

type ProfileType struct {
	Id            int64
	Name          string
	ProfileImage  string
	FollowCount   int64
	FollowerCount int64
}
