package models

import "time"

type Like struct {
	LikeID    string    `json:"like_id"`
	TweetID   string    `json:"tweet_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateLike struct {
	TweetID string `json:"tweet_id"`
	UserID  string `json:"user_id"`
}

type UpdateLike struct {
	LikeID   string `json:"like_id"`
	TweetID  *string `json:"tweet_id,omitempty"`
	UserID   *string `json:"user_id,omitempty"`
}

type LikesResponse struct {
	Likes []Like `json:"likes"`
	Count int    `json:"count"`
}
