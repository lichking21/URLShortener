package redismanager

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func CreateClient() (*redis.Client, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatalln("(FATAL) >> .env file wasn't found")
	}

	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		log.Fatalln("(FATAL) >> redisUrl is empty")
	}

	options, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Fatalln("(FATAL) >> Failed to parse redis url: ", err)
	}

	client := redis.NewClient(options)
	log.Println("(LOG) >> Client was successfully created!")

	return client, nil
}

func ReadFromCache(ctx context.Context, shortCode string, client *redis.Client) (string, error) {

	data, err := client.Get(ctx, shortCode).Result()

	if err != nil {
		return "", err
	}

	return data, nil
}

func WriteToCache(ctx context.Context, shortCode string, longUrl string, client *redis.Client) error {

	if err := client.Set(ctx, shortCode, longUrl, 24*time.Hour).Err(); err != nil {
		return err
	}

	return nil
}
