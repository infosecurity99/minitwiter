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

type retweetRepo struct {
	db    *pgxpool.Pool
	log   logger.ILogger
	redis storage.IRedisStorage
}

func NewRetweetsRepo(db *pgxpool.Pool, log logger.ILogger, redis storage.IRedisStorage) storage.IRetweetsStorage {
	return &retweetRepo{
		db:    db,
		log:   log,
		redis: redis,
	}
}

func (r *retweetRepo) Create(ctx context.Context, retweet models.CreateRetweet) (string, error) {
	id := uuid.New()

	query := `INSERT INTO retweets (retweet_id, original_tweet_id, user_id) VALUES ($1, $2, $3)`
	cmdTag, err := r.db.Exec(ctx, query, id, retweet.OriginalTweetID, retweet.UserID)
	if err != nil {
		r.log.Error("error while inserting retweet data", logger.Error(err))
		return "", err
	}

	if cmdTag.RowsAffected() == 0 {
		r.log.Error("no rows affected while inserting retweet", logger.Error(err))
		return "", fmt.Errorf("no rows affected")
	}

	return id.String(), nil
}

func (r *retweetRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.Retweet, error) {
	retweet := models.Retweet{}
	query := `SELECT retweet_id, original_tweet_id, user_id, created_at FROM retweets WHERE retweet_id = $1`
	err := r.db.QueryRow(ctx, query, key.ID).Scan(&retweet.RetweetID, &retweet.OriginalTweetID, &retweet.UserID, &retweet.CreatedAt)
	if err != nil {
		r.log.Error("error while selecting retweet", logger.Error(err))
		return models.Retweet{}, err
	}

	return retweet, nil
}

func (r *retweetRepo) GetList(ctx context.Context, req models.GetListRequest) (models.RetweetsResponse, error) {
	var (
		retweets = []models.Retweet{}
		count    = 0
		filter   string
		page     = req.Page
		offset   = (page - 1) * req.Limit
	)

	if req.UserID != "" {
		filter += fmt.Sprintf(" AND user_id = '%s'", req.UserID)
	}

	countQuery := `SELECT COUNT(1) FROM retweets WHERE TRUE` + filter
	err := r.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		r.log.Error("error while selecting count", logger.Error(err))
		return models.RetweetsResponse{}, err
	}

	query := `SELECT retweet_id, original_tweet_id, user_id, created_at FROM retweets WHERE TRUE` + filter
	query += ` ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		r.log.Error("error while selecting retweets", logger.Error(err))
		return models.RetweetsResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		retweet := models.Retweet{}
		if err := rows.Scan(&retweet.RetweetID, &retweet.OriginalTweetID, &retweet.UserID, &retweet.CreatedAt); err != nil {
			r.log.Error("error while scanning data", logger.Error(err))
			return models.RetweetsResponse{}, err
		}
		retweets = append(retweets, retweet)
	}

	if err := rows.Err(); err != nil {
		r.log.Error("error while iterating rows", logger.Error(err))
		return models.RetweetsResponse{}, err
	}

	return models.RetweetsResponse{
		Retweets: retweets,
		Count:    count,
	}, nil
}

func (r *retweetRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `DELETE FROM retweets WHERE retweet_id = $1`
	cmdTag, err := r.db.Exec(ctx, query, key.ID)
	if err != nil {
		r.log.Error("error while deleting retweet", logger.Error(err))
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		r.log.Error("no rows affected while deleting retweet", logger.Error(err))
		return fmt.Errorf("no rows affected")
	}

	return nil
}
