package service

import (
	"context"
	"log"
	"urlShortener/base62"
	redismanager "urlShortener/redis_manager"
	"urlShortener/storage"

	"github.com/redis/go-redis/v9"
)

func NewService(repository *storage.PostgresRepo, redisClient *redis.Client) Service {
	return Service{
		repo:  repository,
		redis: redisClient,
	}
}

func (service Service) CreateShortCode(ctx context.Context, url string) (string, error) {

	id, addErr := service.repo.AddOrigUrl(ctx, url)
	if addErr != nil {
		return "", addErr
	}

	shortCode := base62.Encode(id)

	if err := service.repo.UpdateShortCode(ctx, id, shortCode); err != nil {
		return shortCode, err
	}

	return shortCode, nil
}

func (service Service) GetLongUrl(ctx context.Context, shortCode string) (string, error) {

	longUrl, readErr := redismanager.ReadFromCache(ctx, shortCode, service.redis)

	if readErr == nil {
		return longUrl, nil
	}

	if readErr != nil && readErr != redis.Nil {
		log.Println("(WARN) >> Fail to read from Redis: ", readErr)
	}

	url, getErr := service.repo.GetByShortCode(ctx, shortCode)
	if getErr != nil {
		return "", getErr
	}

	if writeErr := redismanager.WriteToCache(ctx, shortCode, url.OrigUrl, service.redis); writeErr != nil {
		log.Println("(WARN) >> Fail to write to Redis: ", writeErr)
	}

	return url.OrigUrl, nil
}
