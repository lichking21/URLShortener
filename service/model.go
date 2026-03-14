package service

import (
	"urlShortener/storage"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	repo  *storage.PostgresRepo
	redis *redis.Client
}
