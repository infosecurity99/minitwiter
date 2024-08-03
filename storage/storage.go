package storage

import (
	"context"
	"test/api/models"
)

type IStorage interface {
	User() IUserStorage
	Tweets() ITweetsStorage
	Followers() IFollowersStorage
	Likes()  ILikesStorage
	ReTweets() IReTweetsStorage
}

type IUserStorage interface {
	Create(context.Context, models.CreateUser) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.User, error)
	GetList(context.Context, models.GetListRequest) (models.UsersResponse, error)
	Update(context.Context, models.UpdateUser) (string, error)
	Delete(context.Context, models.PrimaryKey) error
	GetPassword(context.Context, string) (string, error)
	UpdatePassword(context.Context, models.UpdateUserPassword) error
	GetCustomerCredentialsByLogin(context.Context, string) (models.User, error)
	GetAdminCredentialsByLogin(context.Context, string) (models.User, error)
}



type ITweetsStorage struct {
	Create(context.Context, models.CreateTweet) (string ,error)
	GetByID(context.Context, models.PrimaryKey) (models.Tweet, error)
	GetList(context.Context, models.GetListRequest) (models.TweetsResponse, error)
	Update(context.Context,models.UPdateTweet) (string ,error)
	Delete(context.Context, models.PrimaryKey) error
}

type IFollowersStorage struct {
	Create(context.Context,  models.CreateFollower) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.Followers, error)
	GetList(context.Context, models.GetListRequest) (models.FollowersResponse, error)
	Update(context.Context, models.UpdateFollower) (string, error)
	Delete(context.Context, models.PrimaryKey) error

}

type ILikesStorage  struct {
	Create(context.Context, models.CreateLike) (string, error)
	GetByID(context.Context , models.PrimaryKey) (models.Likes, error)
	GetList(context.Context  , models.GetListRequest) (models.LikesResponse, error)
	Update(context.Context  , models.UpdateLike) (string, error)
	Delete(context.Context , models.PrimaryKey) error
}

type IReTweetsStorage struct {
	Create(context.Context , models.CreateRetweet) (string, error)
	GetById(context.Context  ,models.PrimaryKey) (models.Retweet, error)
	GetList(context.Context  ,models.GetListRequest) (models.RetweetsResponse, error)
	Update(context.Context , models.UpdateRetweet) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}