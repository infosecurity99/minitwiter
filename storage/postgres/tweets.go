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

type tweetRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewTweetRepo(db *pgxpool.Pool, log logger.ILogger) storage.ITweetsStorage {
	return &tweetRepo{
		db:  db,
		log: log,
	}
}

func (t *tweetRepo) Create(ctx context.Context, createTweet models.CreateTweet) (string, error) {
	id := uuid.New()

	query := `
		INSERT INTO tweets (tweet_id, user_id, content, image_url, video_url)
		VALUES ($1, $2, $3, $4, $5)
	`
	cmdTag, err := t.db.Exec(ctx, query, id, createTweet.UserID, createTweet.Content, createTweet.ImageURL, createTweet.VideoURL)
	if err != nil {
		t.log.Error("error while inserting tweet data", logger.Error(err))
		return "", err
	}

	if cmdTag.RowsAffected() == 0 {
		t.log.Error("no rows affected while inserting tweet", logger.Error(err))
		return "", fmt.Errorf("no rows affected")
	}

	return id.String(), nil
}

func (t *tweetRepo) GetByID(ctx context.Context, tweetID models.PrimaryKey) (models.Tweet, error) {
	tweet := models.Tweet{}

	query := `
		SELECT tweet_id, user_id, content, image_url, video_url, created_at, updated_at
		FROM tweets
		WHERE tweet_id = $1
	`
	err := t.db.QueryRow(ctx, query, tweetID).Scan(&tweet.ID, &tweet.UserID, &tweet.Content, &tweet.ImageURL, &tweet.VideoURL, &tweet.CreatedAt, &tweet.UpdatedAt)
	if err != nil {
		t.log.Error("error while scanning tweet", logger.Error(err))
		return models.Tweet{}, err
	}

	return tweet, nil
}

func (t *tweetRepo) GetList(ctx context.Context, request models.GetListRequest) (models.TweetsResponse, error) {
	var (
		tweets = []models.Tweet{}
		count  = 0
		page   = request.Page
		offset = (page - 1) * request.Limit
	)

	countQuery := `
		SELECT COUNT(1)
		FROM tweets
	`

	err := t.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		t.log.Error("error while counting tweets", logger.Error(err))
		return models.TweetsResponse{}, err
	}

	query := `
		SELECT tweet_id, user_id, content, image_url, video_url, created_at, updated_at
		FROM tweets
		ORDER BY created_at DESC LIMIT $1 OFFSET $2
	`
	rows, err := t.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		t.log.Error("error while querying tweets", logger.Error(err))
		return models.TweetsResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		tweet := models.Tweet{}
		if err := rows.Scan(&tweet.ID, &tweet.UserID, &tweet.Content, &tweet.ImageURL, &tweet.VideoURL, &tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
			t.log.Error("error while scanning tweet row", logger.Error(err))
			return models.TweetsResponse{}, err
		}
		tweets = append(tweets, tweet)
	}

	return models.TweetsResponse{
		Tweets: tweets,
		Count:  count,
	}, nil
}

func (t *tweetRepo) Update(ctx context.Context, updateTweet models.UpdateTweet) (string, error) {
	query := `
		UPDATE tweets
		SET content = $1, image_url = $2, video_url = $3, updated_at = NOW()
		WHERE tweet_id = $4
	`
	cmdTag, err := t.db.Exec(ctx, query, updateTweet.Content, updateTweet.ImageURL, updateTweet.VideoURL, updateTweet.ID)
	if err != nil {
		t.log.Error("error while updating tweet data", logger.Error(err))
		return "", err
	}

	if cmdTag.RowsAffected() == 0 {
		t.log.Error("no rows affected while updating tweet", logger.Error(err))
		return "", fmt.Errorf("no rows affected")
	}

	return updateTweet.ID, nil
}

func (t *tweetRepo) Delete(ctx context.Context, tweetID models.PrimaryKey) error {
	query := `DELETE FROM tweets WHERE tweet_id = $1`
	cmdTag, err := t.db.Exec(ctx, query, tweetID)
	if err != nil {
		t.log.Error("error while deleting tweet", logger.Error(err))
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		t.log.Error("no rows affected while deleting tweet", logger.Error(err))
		return fmt.Errorf("no rows affected")
	}

	return nil
}
