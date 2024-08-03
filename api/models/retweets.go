package models

import "time"

type Retweet struct {
	RetweetID         string    `json:"retweet_id"`
	OriginalTweetID   string    `json:"original_tweet_id"`
	UserID            string    `json:"user_id"`
	CreatedAt         time.Time `json:"created_at"`
}

type CreateRetweet struct {
	OriginalTweetID string `json:"original_tweet_id"`
	UserID          string `json:"user_id"`
}

type UpdateRetweet struct {
	RetweetID        string `json:"retweet_id"`
	OriginalTweetID  *string `json:"original_tweet_id,omitempty"`
	UserID           *string `json:"user_id,omitempty"`
}

type RetweetsResponse struct {
	Retweets []Retweet `json:"retweets"`
	Count    int      `json:"count"`
}
