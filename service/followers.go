package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type followerService struct {
	storage storage.IStorage
	log     logger.ILogger
	redis   storage.IRedisStorage
}

func NewFollowerService(storage storage.IStorage, log logger.ILogger, redis storage.IRedisStorage) followerService {
	return followerService{storage: storage, log: log, redis: redis}
}

func (f followerService) Follow(ctx context.Context, follow models.Follow) error {
	f.log.Info("follower create service layer", logger.Any("follow", follow))

	err := f.storage.Follower().Create(ctx, follow)
	if err != nil {
		f.log.Error("error in service layer while creating follow relationship", logger.Error(err))
		return err
	}

	return nil
}

func (f followerService) Unfollow(ctx context.Context, unfollow models.Follow) error {
	f.log.Info("follower delete service layer", logger.Any("unfollow", unfollow))

	err := f.storage.Follower().Delete(ctx, unfollow)
	if err != nil {
		f.log.Error("error in service layer while deleting follow relationship", logger.Error(err))
		return err
	}

	return nil
}

func (f followerService) GetFollowers(ctx context.Context, userID string) ([]models.User, error) {
	f.log.Info("get followers service layer", logger.Any("userID", userID))

	followers, err := f.storage.Follower().GetFollowersByUserID(ctx, userID)
	if err != nil {
		f.log.Error("error in service layer while getting followers", logger.Error(err))
		return nil, err
	}

	return followers, nil
}

func (f followerService) GetFollowing(ctx context.Context, userID string) ([]models.User, error) {
	f.log.Info("get following service layer", logger.Any("userID", userID))

	following, err := f.storage.Follower().GetFollowingByUserID(ctx, userID)
	if err != nil {
		f.log.Error("error in service layer while getting following", logger.Error(err))
		return nil, err
	}

	return following, nil
}
