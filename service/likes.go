package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type likesService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewlikesService(storage storage.IStorage, log logger.ILogger) likesService {
	return likesService{storage: storage, log: log}
}

func (t likesService) Create(ctx context.Context, tweet models.CreateTweet) (models.Tweet, error) {
	t.log.Info("likesService create service layer", logger.Any("likesService", tweet))

	id, err := t.storage.Tweets().Create(ctx, tweet)
	if err != nil {
		t.log.Error("error in service layer while creating tweet", logger.Error(err))
		return models.Tweet{}, err
	}

	createdTweet, err := t.storage.Tweets().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		t.log.Error("error in service layer while getting tweet by id", logger.Error(err))
		return models.Tweet{}, err
	}

	return createdTweet, nil
}

func (t likesService) Get(ctx context.Context, id string) (models.Tweet, error) {
	tweet, err := t.storage.Tweets().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		t.log.Error("error in service layer while getting tweet by id", logger.Error(err))
		return models.Tweet{}, err
	}

	return tweet, nil
}





func (t likesService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := t.storage.Tweets().Delete(ctx, key)
	return err
}
