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

type retweetsRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewReTweetsRepo(db *pgxpool.Pool, log logger.ILogger) storage.IRetweetsStorage {
	return &retweetsRepo{
		db:  db,
		log: log,
	}
}

func (r *retweetsRepo) Create(ctx context.Context, retweet models.CreateRetweet) (string, error) {
	id := uuid.New()

	query := `INSERT INTO retweets (retweet_id, tweet_id, user_id) VALUES ($1, $2, $3)`
	cmdTag, err := r.db.Exec(ctx, query, id, retweet.OriginalTweetID, retweet.UserID)
	if err != nil {
		r.log.Error("Error while inserting retweet data", logger.Error(err))
		return "", err
	}

	if cmdTag.RowsAffected() == 0 {
		err := fmt.Errorf("No rows affected while inserting retweet")
		r.log.Error(err.Error())
		return "", err
	}

	return id.String(), nil
}

func (r *retweetsRepo) Delete(ctx context.Context, retweetID models.PrimaryKey) error {
	query := `DELETE FROM retweets WHERE retweet_id = $1`
	cmdTag, err := r.db.Exec(ctx, query, retweetID)
	if err != nil {
		r.log.Error("Error while deleting retweet", logger.Error(err))
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		err := fmt.Errorf("No rows affected while deleting retweet")
		r.log.Error(err.Error())
		return err
	}

	return nil
}
