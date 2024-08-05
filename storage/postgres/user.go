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

type userRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewUserRepo(db *pgxpool.Pool, log logger.ILogger) storage.IUserStorage {
	return &userRepo{
		db:  db,
		log: log,
	}
}

func (u *userRepo) Create(ctx context.Context, createUser models.CreateUser) (string, error) {
	uid := uuid.New()

	query := `
		INSERT INTO users (user_id, username, password_hash, name, bio, profile_picture)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	cmdTag, err := u.db.Exec(ctx, query, uid, createUser.Username, createUser.Password, createUser.Name, createUser.Bio, createUser.ProfilePicture)
	if err != nil {
		u.log.Error("error while inserting user data", logger.Error(err))
		return "", err
	}

	if cmdTag.RowsAffected() == 0 {
		u.log.Error("no rows affected while inserting user", logger.Error(err))
		return "", fmt.Errorf("no rows affected")
	}

	return uid.String(), nil
}

func (u *userRepo) GetByID(ctx context.Context, pKey models.PrimaryKey) (models.User, error) {
	user := models.User{}

	query := `
		SELECT user_id, username, password_hash, name, bio, profile_picture, created_at, updated_at
		FROM users
		WHERE user_id = $1
	`
	err := u.db.QueryRow(ctx, query, pKey.ID).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Name, &user.Bio, &user.ProfilePicture, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		u.log.Error("error while scanning user", logger.Error(err))
		return models.User{}, err
	}

	return user, nil
}

func (u *userRepo) GetList(ctx context.Context, request models.GetListRequest) (models.UsersResponse, error) {
	var (
		users  = []models.User{}
		count  = 0
		page   = request.Page
		offset = (page - 1) * request.Limit
		search = request.Search
	)

	countQuery := `
		SELECT COUNT(1)
		FROM users
	`

	if search != "" {
		countQuery += ` WHERE username ILIKE $1 OR name ILIKE $2`
	}

	args := []interface{}{}
	if search != "" {
		args = append(args, search, search)
	}

	err := u.db.QueryRow(ctx, countQuery, args...).Scan(&count)
	if err != nil {
		u.log.Error("error while counting users", logger.Error(err))
		return models.UsersResponse{}, err
	}

	query := `
		SELECT user_id, username, password_hash, name, bio, profile_picture, created_at, updated_at
		FROM users
	`

	if search != "" {
		query += ` WHERE username ILIKE $1 OR name ILIKE $2`
	}

	query += ` ORDER BY created_at DESC LIMIT $3 OFFSET $4`
	args = append(args, request.Limit, offset)

	rows, err := u.db.Query(ctx, query, args...)
	if err != nil {
		u.log.Error("error while querying users", logger.Error(err))
		return models.UsersResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Name, &user.Bio, &user.ProfilePicture, &user.CreatedAt, &user.UpdatedAt); err != nil {
			u.log.Error("error while scanning user row", logger.Error(err))
			return models.UsersResponse{}, err
		}
		users = append(users, user)
	}

	return models.UsersResponse{
		Users: users,
		Count: count,
	}, nil
}

func (u *userRepo) Update(ctx context.Context, request models.UpdateUser) (models.User, error) {
	query := `
		UPDATE users
		SET  name = $1, bio = $2, profile_picture = $3, updated_at = NOW()
		WHERE user_id = $4
	`
	cmdTag, err := u.db.Exec(ctx, query, request.Name, request.Bio, request.ProfilePicture, request.ID)
	if err != nil {
		u.log.Error("error while updating user data", logger.Error(err))
		return models.User{}, err
	}

	if cmdTag.RowsAffected() == 0 {
		u.log.Error("no rows affected while updating user", logger.Error(err))
		return models.User{}, fmt.Errorf("no rows affected")
	}

	return models.User{}, nil
}

func (u *userRepo) Delete(ctx context.Context, request models.PrimaryKey) error {
	query := `UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE user_id = $1`
	cmdTag, err := u.db.Exec(ctx, query, request.ID)
	if err != nil {
		u.log.Error("error while deleting user", logger.Error(err))
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		u.log.Error("no rows affected while deleting user", logger.Error(err))
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func (u *userRepo) GetPassword(ctx context.Context, id models.PrimaryKey) (string, error) {
	var password string
	query := `SELECT password_hash FROM users WHERE user_id = $1`
	if err := u.db.QueryRow(ctx, query, id).Scan(&password); err != nil {
		u.log.Error("error while retrieving user password", logger.Error(err))
		return "", err
	}

	return password, nil
}

func (u *userRepo) UpdatePassword(ctx context.Context, request models.UpdateUserPassword) error {
	query := `UPDATE users SET password_hash = $1, updated_at = NOW() WHERE user_id = $2`
	cmdTag, err := u.db.Exec(ctx, query, request.NewPassword, request.ID)
	if err != nil {
		u.log.Error("error while updating user password", logger.Error(err))
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		u.log.Error("no rows affected while updating user password", logger.Error(err))
		return fmt.Errorf("no rows affected")
	}

	return nil
}



func (u *userRepo) GetUserCredentialsByLogin(ctx context.Context, login string) (models.User, error) {
	user := models.User{}
	query := `SELECT user_id, password_hash FROM users WHERE username = $1`
	if err := u.db.QueryRow(ctx, query, login).Scan(&user.ID, &user.PasswordHash); err != nil {
		u.log.Error("error while retrieving admin credentials by login", logger.Error(err))
		return models.User{}, err
	}

	return user, nil
}

