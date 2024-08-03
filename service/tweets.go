package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type tweetService struct {
	storage storage.IStorage
	log     logger.ILogger
	redis   storage.IRedisStorage
}

func NewTweetService(storage storage.IStorage, log logger.ILogger, redis storage.IRedisStorage) tweetService {
	return tweetService{storage: storage, log: log, redis: redis}
}

func (t tweetService) Create(ctx context.Context, tweet models.CreateTweet) (models.Tweet, error) {
	t.log.Info("tweet create service layer", logger.Any("tweet", tweet))

	id, err := t.storage.Tweet().Create(ctx, tweet)
	if err != nil {
		t.log.Error("error in service layer while creating tweet", logger.Error(err))
		return models.Tweet{}, err
	}

	createdTweet, err := t.storage.Tweet().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		t.log.Error("error in service layer while getting tweet by id", logger.Error(err))
		return models.Tweet{}, err
	}

	return createdTweet, nil
}

func (t tweetService) Get(ctx context.Context, id string) (models.Tweet, error) {
	tweet, err := t.storage.Tweet().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		t.log.Error("error in service layer while getting tweet by id", logger.Error(err))
		return models.Tweet{}, err
	}

	return tweet, nil
}

func (t tweetService) GetList(ctx context.Context, request models.GetListRequest) (models.TweetResponse, error) {
	t.log.Info("tweet get list service layer", logger.Any("request", request))

	tweets, err := t.storage.Tweet().GetList(ctx, request)
	if err != nil {
		t.log.Error("error in service layer while getting list of tweets", logger.Error(err))
		return models.TweetResponse{}, err
	}

	return tweets, nil
}

func (t tweetService) Update(ctx context.Context, tweet models.UpdateTweet) (models.Tweet, error) {
	id, err := t.storage.Tweet().Update(ctx, tweet)
	if err != nil {
		t.log.Error("error in service layer while updating tweet", logger.Error(err))
		return models.Tweet{}, err
	}

	updatedTweet, err := t.storage.Tweet().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		t.log.Error("error in service layer while getting updated tweet by id", logger.Error(err))
		return models.Tweet{}, err
	}

	return updatedTweet, nil
}

func (t tweetService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := t.storage.Tweet().Delete(ctx, key)
	return err
}