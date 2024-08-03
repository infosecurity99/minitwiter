package service

import (
	"test/pkg/logger"
	"test/storage"
)

type IServiceManager interface {
	User() userService
	Tweets() tweetsService
	Followers() followersService
	Likes() likesService
	Retweets() retweetsSerive

	AuthService() authService
	RedisService() redisService
}

type Service struct {
	userService      userService
	tweetsService    tweetsService
	followersService followersService
	likesService     likesService
	retweets         retweetsSerice

	authService  authService
	redisService redisService
}

func New(storage storage.IStorage, log logger.ILogger, redis storage.IRedisStorage) Service {
	services := Service{}

	services.userService = NewUserService(storage, log, redis)
	services.authService = NewAuthService(storage, log, redis)
	services.redisService = NewRedisService(storage, log, redis)
	return services
}

func (s Service) User() userService {
	return s.userService
}

func (s Service) Tweets() tweetsService {
	return s.tweetService
}

func (s Service) Followers() followerService {
	return s.followersSerive
}

func (s Service) Likes() likesService {
	return s.likesService
}

func (s Service) Retweets() retweetsSerivce {
	return s.retweets
}

func (s Service) AuthService() authService {
	return s.authService
}

func (s Service) RedisService() redisService {
	return s.redisService
}
