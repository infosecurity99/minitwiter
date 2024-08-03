package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type likeService struct {
	storage storage.IStorage
	log     logger.ILogger
	redis   storage.IRedisStorage
}

func NewLikeService(storage storage.IStorage, log logger.ILogger, redis storage.IRedisStorage) likeService {
	return likeService{storage: storage, log: log, redis: redis}
}

func (l likeService) Like(ctx context.Context, like models.Like) error {
	l.log.Info("like create service layer", logger.Any("like", like))

	err := l.storage.Like().Create(ctx, like)
	if err != nil {
		l.log.Error("error in service layer while creating like", logger.Error(err))
		return err
	}

	return nil
}

func (l likeService) Unlike(ctx context.Context, unlike models.Like) error {
	l.log.Info("like delete service layer", logger.Any("unlike", unlike))

	err := l.storage.Like().Delete(ctx, unlike)
	if err != nil {
		l.log.Error("error in service layer while deleting like", logger.Error(err))
		return err
	}

	return nil
}

func (l likeService) GetLikesByTweet(ctx context.Context, tweetID string) ([]models.User, error) {
	l.log.Info("get likes by tweet service layer", logger.Any("tweetID", tweetID))

	users, err := l.storage.Like().GetLikesByTweetID(ctx, tweetID)
	if err != nil {
		l.log.Error("error in service layer while getting likes by tweet", logger.Error(err))
		return nil, err
	}

	return users, nil
}

func (l likeService) GetLikedTweetsByUser(ctx context.Context, userID string) ([]models.Tweet, error) {
	l.log.Info("get liked tweets by user service layer", logger.Any("userID", userID))

	tweets, err := l.storage.Like().GetLikedTweetsByUserID(ctx, userID)
	if err != nil {
		l.log.Error("error in service layer while getting liked tweets by user", logger.Error(err))
		return nil, err
	}

	return tweets, nil
}
