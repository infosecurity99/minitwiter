package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type followersService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewfollowersService(storage storage.IStorage, log logger.ILogger) followersService {
	return followersService{storage: storage, log: log}
}

func (f followersService) Create(ctx context.Context, follower models.CreateFollower) (models.Follower, error) {
    f.log.Info("follower create service layer", logger.Any("follower", follower))

    id, err := f.storage.Followers().Create(ctx, follower)
    if err != nil {
        f.log.Error("error in service layer while creating follower", logger.Error(err))
        return models.Follower{}, err
    }

    createdFollower, err := f.storage.Followers().GetByID(ctx, models.PrimaryKey{ID: id})
    if err != nil {
        f.log.Error("error in service layer while getting follower by id", logger.Error(err))
        return models.Follower{}, err
    }

    return createdFollower, nil
}

func (f followersService) Get(ctx context.Context, id string) (models.Follower, error) {
    follower, err := f.storage.Followers().GetByID(ctx, models.PrimaryKey{ID: id})
    if err != nil {
        f.log.Error("error in service layer while getting follower by id", logger.Error(err))
        return models.Follower{}, err
    }

    return follower, nil
}


func (f followersService) GetList(ctx context.Context, request models.GetListRequest) (models.FollowersResponse, error) {
    f.log.Info("follower get list service layer", logger.Any("request", request))

    followers, err := f.storage.Followers().GetList(ctx, request)
    if err != nil {
        f.log.Error("error in service layer while getting list of followers", logger.Error(err))
        return models.FollowersResponse{}, err
    }

    return followers, nil
}


func (f followersService) Delete(ctx context.Context, key models.PrimaryKey) error {
    err := f.storage.Followers().Delete(ctx, key)
    return err
}
