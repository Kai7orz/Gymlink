package dto

type UserType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type UserCreateType struct {
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

type ProfileType struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	ProfileImage  string `json:"profile_image"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
}

type UserFollowType struct {
	FollowerId int64 `json:"follower_id" db:"follower_id"`
	FollowedId int64 `json:"followed_id" db:"followed_id"`
}
