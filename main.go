package main

import (
	"context"
	"log"
	"urlShortener/base62"
	dbmanager "urlShortener/db_manager"
	"urlShortener/storage"
)

func main() {

	dbPool, err := dbmanager.ConnectToDb()
	if err != nil {
		log.Fatalln("(FATAL) >> Failed to connect to db: ", err)
	}

	ctx := context.Background()
	// redismanager.CreateClient()

	urlString := "https://github.com/lichking21"

	repo := storage.NewPostgresRepo(dbPool)

	id, _ := repo.AddOrigUrl(ctx, urlString)
	shortCode := base62.Encode(id)
	repo.UpdateShortCode(ctx, id, shortCode)

	url, _ := repo.GetByShortCode(ctx, shortCode)

	storage.PrintUrl(url)
}
