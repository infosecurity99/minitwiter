package service

import (
	"context"
	"errors"
	"test/api/models"
	"test/pkg/check"
	"test/pkg/logger"
	"test/pkg/security"
	"test/storage"

	"github.com/jackc/pgx/v5"
)

type userService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewuserService(storage storage.IStorage, log logger.ILogger) storage.IUserStorage {
	return &userService{
		storage: storage,
		log:     log,
	}
}

func (u *userService) Create(ctx context.Context, createUser models.CreateUser) (string, error) {
	u.log.Info("User create service layer", logger.Any("createUser", createUser))

	password, err := security.HashPassword(createUser.Password)
	if err != nil {
		u.log.Error("Error while hashing password", logger.Error(err))
		return "", err
	}
	createUser.Password = password

	id, err := u.storage.User().Create(ctx, createUser)
	if err != nil {
		u.log.Error("Error while creating user", logger.Error(err))
		return "", err
	}

	user, err := u.storage.User().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		u.log.Error("Error while retrieving created user", logger.Error(err))
		return "", err
	}

	return user.ID, nil
}

func (u *userService) GetByID(ctx context.Context, id models.PrimaryKey) (models.User, error) {
	return u.storage.User().GetByID(ctx, id)
}

func (u *userService) GetList(ctx context.Context, request models.GetListRequest) (models.UsersResponse, error) {
	u.log.Info("Get user list service layer", logger.Any("request", request))

	usersResponse, err := u.storage.User().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			u.log.Error("Error while getting users list", logger.Error(err))
		}
		return models.UsersResponse{}, err
	}

	return usersResponse, nil
}

func (u *userService) Update(ctx context.Context, updateUser models.UpdateUser) (models.User, error) {
	if _, err := u.storage.User().Update(ctx, updateUser); err != nil {
		u.log.Error("Error while updating user", logger.Error(err))
		return models.User{}, err
	}

	user, err := u.storage.User().GetByID(ctx, models.PrimaryKey{ID: updateUser.ID})
	if err != nil {
		u.log.Error("Error while getting user after update", logger.Error(err))
		return models.User{}, err
	}

	return user, nil
}

func (u *userService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := u.storage.User().Delete(ctx, key)
	if err != nil {
		u.log.Error("Error while deleting user", logger.Error(err))
	}
	return err
}

func (u *userService) UpdatePassword(ctx context.Context, request models.UpdateUserPassword) error {
	oldPasswordHash, err := u.storage.User().GetPassword(ctx, models.PrimaryKey{})
	if err != nil {
		u.log.Error("Error while retrieving current password hash", logger.Error(err))
		return err
	}

	if err := security.CompareHashAndPassword(request.OldPassword, oldPasswordHash); err != nil {
		u.log.Error("Old password did not match", logger.Error(err))
		return errors.New("old password did not match")
	}

	if err := check.ValidatePassword(request.NewPassword); err != nil {
		u.log.Error("New password is weak", logger.Error(err))
		return err
	}

	newPasswordHash, err := security.HashPassword(request.NewPassword)
	if err != nil {
		u.log.Error("Error while hashing new password", logger.Error(err))
		return err
	}

	if err := u.storage.User().UpdatePassword(ctx, models.UpdateUserPassword{
		ID:          request.ID,
		NewPassword: newPasswordHash,
	}); err != nil {
		u.log.Error("Error while updating password", logger.Error(err))
		return err
	}

	return nil
}

func (u *userService) GetAdminCredentialsByLogin(ctx context.Context, login string) (models.User, error) {
	user, err := u.storage.User().GetAdminCredentialsByLogin(ctx, login)
	if err != nil {
		u.log.Error("Error while retrieving admin credentials by login", logger.Error(err))
		return models.User{}, err
	}
	return user, nil
}

func (u *userService) GetCustomerCredentialsByLogin(ctx context.Context, login string) (models.User, error) {
	user, err := u.storage.User().GetCustomerCredentialsByLogin(ctx, login)
	if err != nil {
		u.log.Error("Error while retrieving customer credentials by login", logger.Error(err))
		return models.User{}, err
	}
	return user, nil
}
func (u *userService) GetPassword(ctx context.Context, id models.PrimaryKey) (string, error) {
    user, err := u.storage.User().GetByID(ctx, id)
    if err != nil {
        u.log.Error("Error while retrieving user", logger.Error(err))
        return "", err
    }
    return user.PasswordHash, nil // Ensure `Password` field is accessible
}