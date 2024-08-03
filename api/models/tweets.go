package models

import "time"

type Tweet struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	ImageURL  *string   `json:"image_url,omitempty"`
	VideoURL  *string   `json:"video_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTweet struct {
	UserID   string  `json:"user_id"`
	Content  string  `json:"content"`
	ImageURL *string `json:"image_url,omitempty"`
	VideoURL *string `json:"video_url,omitempty"`
}

type UpdateTweet struct {
	ID       string  `json:"id"`
	Content  *string `json:"content,omitempty"`
	ImageURL *string `json:"image_url,omitempty"`
	VideoURL *string `json:"video_url,omitempty"`
}

type TweetsResponse struct {
	Tweets []Tweet `json:"tweets"`
	Count  int     `json:"count"`
}
