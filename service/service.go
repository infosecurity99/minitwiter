package service

import (
	"test/pkg/logger"
	"test/storage"
)

type IServiceManager interface {
	User() userService
	Tweets() tweetService
	Followers() followersService
	Likes() likesService

	AuthService() authService
}

type Service struct {
	userService      userService
	tweetsService    tweetService
	followersService followersService
	likesService     likesService

	authService authService
}

func New(storage storage.IStorage, log logger.ILogger) Service {
	services := Service{}
	services.tweetsService = NewTweetService(storage, log)
	services.followersService = NewfollowersService(storage, log)
	services.likesService = NewlikesService(storage, log)

	services.userService= NewuserService(storage, log) 
	services.authService = NewAuthService(storage, log)
	return services
}

func (s Service) User() userService {
	return s.userService
}

func (s Service) Tweets() tweetService {
	return s.tweetsService
}

func (s Service) Followers() followersService {
	return s.followersService
}

func (s Service) Likes() likesService {
	return s.likesService
}

func (s Service) AuthService() authService {
	return s.authService
}
