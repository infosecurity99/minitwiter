package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type retweetsService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewretweetsSerice(storage storage.IStorage, log logger.ILogger) retweetsService {
	return retweetsService{storage: storage, log: log}
}

func (r retweetsService) Create(ctx context.Context, retweet models.CreateRetweet) (string, error) {
	r.log.Info("retweetsService create service layer", logger.Any("retweet", retweet))

	id, err := r.storage.Retweets().Create(ctx, retweet)
	if err != nil {
		r.log.Error("error in service layer while creating retweet", logger.Error(err))
		return "", err
	}

	return id, nil
}

func (r retweetsService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := r.storage.Retweets().Delete(ctx, key)
	return err
}