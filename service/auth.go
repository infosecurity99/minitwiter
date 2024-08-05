package service

import (
	"context"
	"test/api/models"
	"test/pkg/jwt"
	"test/pkg/logger"
	"test/pkg/security"
	"test/storage"
)

type authService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewAuthService(storage storage.IStorage, log logger.ILogger) authService {
	return authService{
		storage: storage,
		log:     log,
	}
}

func (a authService) UserLogin(ctx context.Context, loginRequest models.UserLoginRequest) (models.UserLoginResponse, error) {
	admin, err := a.storage.User().GetUserCredentialsByLogin(ctx, loginRequest.Login)
	if err != nil {
		a.log.Error("error is while getting user", logger.Error(err))
		return models.UserLoginResponse{}, err
	}

	if err := security.CompareHashAndPassword(admin.PasswordHash, loginRequest.Password); err != nil {
		a.log.Error("password is incorrect", logger.Error(err))
		return models.UserLoginResponse{}, err
	}

	m := make(map[interface{}]interface{})
	m["user_id"] = admin.ID
	m["user_role"] = "admin"

	accessToken, refreshToken, err := jwt.GenerateJWT(m)
	if err != nil {
		a.log.Error("error while generate jwt token", logger.Error(err))
		return models.UserLoginResponse{}, err
	}

	return models.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
