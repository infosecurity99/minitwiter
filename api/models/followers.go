package models

import "time"

type Follower struct {
	FollowerID     string    `json:"follower_id"`
	UserID         string    `json:"user_id"`
	FollowerUserID string    `json:"follower_user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type CreateFollower struct {
	UserID         string `json:"user_id"`
	FollowerUserID string `json:"follower_user_id"`
}

type UpdateFollower struct {
	FollowerID     string  `json:"follower_id"`
	UserID         *string `json:"user_id,omitempty"`
	FollowerUserID *string `json:"follower_user_id,omitempty"`
}

type FollowersResponse struct {
	Followers []Follower `json:"followers"`
	Count     int        `json:"count"`
}
