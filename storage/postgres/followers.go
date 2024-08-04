package postgres

import (
	"context"
	"errors"
	"fmt"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type followerRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewFollowersRepo(db *pgxpool.Pool, log logger.ILogger) storage.IFollowersStorage {
	return &followerRepo{
		db:  db,
		log: log,
	}
}

func (b *followerRepo) Create(ctx context.Context, follower models.CreateFollower) (string, error) {
	if follower.UserID == follower.FollowerUserID {
		err := errors.New("users cannot follow themselves")
		b.log.Error("validation error", logger.Error(err))
		return "", err
	}

	id := uuid.New()

	query := `INSERT INTO followers (follower_id, user_id, follower_user_id) VALUES ($1, $2, $3)`
	if rowsAffected, err := b.db.Exec(ctx, query, id, follower.UserID, follower.FollowerUserID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			b.log.Error("no rows affected while inserting follower", logger.Error(err))
			return "", err
		}
		b.log.Error("error while inserting follower data", logger.Error(err))
		return "", err
	}

	return id.String(), nil
}

func (b *followerRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.Follower, error) {
	follower := models.Follower{}
	query := `SELECT follower_id, user_id, follower_user_id, created_at FROM followers WHERE follower_id = $1`
	if err := b.db.QueryRow(ctx, query, key.ID).Scan(&follower.FollowerID, &follower.UserID, &follower.FollowerUserID, &follower.CreatedAt); err != nil {
		b.log.Error("error while selecting follower", logger.Error(err))
		return models.Follower{}, err
	}

	return follower, nil
}

func (b *followerRepo) GetList(ctx context.Context, req models.GetListRequest) (models.FollowersResponse, error) {
	var (
		followers = []models.Follower{}
		count     = 0
		query     string
		filter    string
		page      = req.Page
		offset    = (page - 1) * req.Limit
	)

	if req.UserID != "" {
		filter += fmt.Sprintf(" AND user_id = '%s'", req.UserID)
	}

	countQuery := `SELECT COUNT(1) FROM followers WHERE TRUE ` + filter
	if err := b.db.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		b.log.Error("error while selecting count", logger.Error(err))
		return models.FollowersResponse{}, err
	}

	query = `SELECT follower_id, user_id, follower_user_id, created_at FROM followers WHERE TRUE ` + filter
	query += ` ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := b.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		b.log.Error("error while selecting followers", logger.Error(err))
		return models.FollowersResponse{}, err
	}

	for rows.Next() {
		follower := models.Follower{}
		if err = rows.Scan(&follower.FollowerID, &follower.UserID, &follower.FollowerUserID, &follower.CreatedAt); err != nil {
			b.log.Error("error while scanning data", logger.Error(err))
			return models.FollowersResponse{}, err
		}
		followers = append(followers, follower)
	}

	return models.FollowersResponse{
		Followers: followers,
		Count:     count,
	}, nil
}

func (b *followerRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `DELETE FROM followers WHERE follower_id = $1`
	if rowsAffected, err := b.db.Exec(ctx, query, key.ID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			b.log.Error("error while deleting follower", logger.Error(err))
			return err
		}
		return err
	}
	return nil
}
