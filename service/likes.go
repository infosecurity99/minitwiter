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

func (l likesService) Create(ctx context.Context, like models.CreateLike) (models.Like, error) {
	l.log.Info("likesService create service layer", logger.Any("like", like))

	id, err := l.storage.Likes().Create(ctx, like)
	if err != nil {
		l.log.Error("error in service layer while creating like", logger.Error(err))
		return models.Like{}, err
	}

	createdLike, err := l.storage.Likes().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		l.log.Error("error in service layer while getting like by id", logger.Error(err))
		return models.Like{}, err
	}

	return createdLike, nil
}

func (l likesService) Get(ctx context.Context, id string) (models.Like, error) {
	like, err := l.storage.Likes().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		l.log.Error("error in service layer while getting like by id", logger.Error(err))
		return models.Like{}, err
	}

	return like, nil
}

func (l likesService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := l.storage.Likes().Delete(ctx, key)
	return err
}
