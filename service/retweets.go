package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type retweetService struct {
	storage storage.IStorage
	log     logger.ILogger
	redis   storage.IRedisStorage
}

func NewRetweetService(storage storage.IStorage, log logger.ILogger, redis storage.IRedisStorage) retweetService {
	return retweetService{storage: storage, log: log, redis: redis}
}

func (r retweetService) Retweet(ctx context.Context, retweet models.Retweet) error {
	r.log.Info("retweet create service layer", logger.Any("retweet", retweet))

	err := r.storage.Retweet().Create(ctx, retweet)
	if err != nil {
		r.log.Error("error in service layer while creating retweet", logger.Error(err))
		return err
	}

	return nil
}

func (r retweetService) RemoveRetweet(ctx context.Context, retweet models.Retweet) error {
	r.log.Info("retweet delete service layer", logger.Any("retweet", retweet))

	err := r.storage.Retweet().Delete(ctx, retweet)
	if err != nil {
		r.log.Error("error in service layer while deleting retweet", logger.Error(err))
		return err
	}

	return nil
}

func (r retweetService) GetRetweetsByTweet(ctx context.Context, tweetID string) ([]models.User, error) {
	r.log.Info("get retweets by tweet service layer", logger.Any("tweetID", tweetID))

	users, err := r.storage.Retweet().GetRetweetsByTweetID(ctx, tweetID)
	if err != nil {
		r.log.Error("error in service layer while getting retweets by tweet", logger.Error(err))
		return nil, err
	}

	return users, nil
}

func (r retweetService) GetRetweetedTweetsByUser(ctx context.Context, userID string) ([]models.Tweet, error) {
	r.log.Info("get retweeted tweets by user service layer", logger.Any("userID", userID))

	tweets, err := r.storage.Retweet().GetRetweetedTweetsByUserID(ctx, userID)
	if err != nil {
		r.log.Error("error in service layer while getting retweeted tweets by user", logger.Error(err))
		return nil, err
	}

	return tweets, nil
}
