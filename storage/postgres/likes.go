package postgres

import (
	"context"
	"fmt"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type likeRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewLikesRepo(db *pgxpool.Pool, log logger.ILogger) storage.ILikesStorage {
	return &likeRepo{
		db:  db,
		log: log,
	}
}

func (l *likeRepo) Create(ctx context.Context, like models.CreateLike) (string, error) {
	id := uuid.New()

	query := `INSERT INTO likes (like_id, tweet_id, user_id) VALUES ($1, $2, $3)`
	cmdTag, err := l.db.Exec(ctx, query, id, like.TweetID, like.UserID)
	if err != nil {
		l.log.Error("Error while inserting like data", logger.Error(err))
		return "", err
	}

	if cmdTag.RowsAffected() == 0 {
		err := fmt.Errorf("No rows affected while inserting like")
		l.log.Error(err.Error())
		return "", err
	}

	return id.String(), nil
}

func (l *likeRepo) GetByID(ctx context.Context, likeID models.PrimaryKey) (models.Like, error) {
	var like models.Like
	query := `SELECT like_id, tweet_id, user_id, created_at FROM likes WHERE like_id = $1`
	err := l.db.QueryRow(ctx, query, likeID).Scan(&like.LikeID, &like.TweetID, &like.UserID, &like.CreatedAt)
	if err != nil {
		l.log.Error("Error while selecting like", logger.Error(err))
		return models.Like{}, err
	}

	return like, nil
}

func (l *likeRepo) Delete(ctx context.Context, likeID models.PrimaryKey) error {
	query := `DELETE FROM likes WHERE like_id = $1`
	cmdTag, err := l.db.Exec(ctx, query, likeID)
	if err != nil {
		l.log.Error("Error while deleting like", logger.Error(err))
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		err := fmt.Errorf("No rows affected while deleting like")
		l.log.Error(err.Error())
		return err
	}

	return nil
}
