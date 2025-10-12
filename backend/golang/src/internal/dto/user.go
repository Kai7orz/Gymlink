package dto

type UserType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type UserCreateType struct {
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

type UserFollowType struct {
	FollowerId int64 `json:"follower_id" db:"follower_id"`
	FollowedId int64 `json:"followed_id" db:"followed_id"`
}
