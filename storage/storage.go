package storage

import (
	"context"
	"test/api/models"
)

type IStorage interface {
	Close()

	User() IUserStorage
	Tweets() ITweetsStorage
	Followers() IFollowersStorage
	Likes() ILikesStorage
	
}

type IUserStorage interface {
	Create(ctx context.Context, createUser models.CreateUser) (string, error)
	GetByID(ctx context.Context, id models.PrimaryKey) (models.User, error)
	GetList(ctx context.Context, request models.GetListRequest) (models.UsersResponse, error)
	Update(ctx context.Context, updateUser models.UpdateUser) (models.User, error)
	Delete(ctx context.Context, key models.PrimaryKey) error
	UpdatePassword(ctx context.Context, request models.UpdateUserPassword) error
	GetAdminCredentialsByLogin(ctx context.Context, login string) (models.User, error)
	GetCustomerCredentialsByLogin(ctx context.Context, login string) (models.User, error)
	GetPassword(ctx context.Context, id models.PrimaryKey) (string, error)
}

type ITweetsStorage interface {
	Create(context.Context, models.CreateTweet) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.Tweet, error)
	GetList(context.Context, models.GetListRequest) (models.TweetsResponse, error)
	Update(context.Context, models.UpdateTweet) (string, error)
	Delete(context.Context, models.PrimaryKey) error
}

type IFollowersStorage interface {
	Create(context.Context, models.CreateFollower) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.Follower, error)
	GetList(context.Context, models.GetListRequest) (models.FollowersResponse, error)
	Delete(context.Context, models.PrimaryKey) error
}

type ILikesStorage interface {
	Create(context.Context, models.CreateLike) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.Like, error)
	Delete(context.Context, models.PrimaryKey) error
}

