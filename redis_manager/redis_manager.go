package redismanager

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func CreateClient() *redis.Client {

	if err := godotenv.Load(); err != nil {
		log.Fatalln("(FATAL) >> .env file wasn't found")
		return nil
	}

	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		log.Fatalln("(FATAL) >> redisUrl is empty")
		return nil
	}

	options, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Fatalln("(FATAL) >> Failed to parse redis url: ", err)
	}

	client := redis.NewClient(options)
	log.Println("(LOG) >> Client was successfully created!")

	return client
}
