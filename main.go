package main

import (
	"log"
	"net/http"
	dbmanager "urlShortener/db_manager"
	"urlShortener/handler"
	redismanager "urlShortener/redis_manager"
	"urlShortener/service"
	"urlShortener/storage"
)

func main() {

	dbPool, dbErr := dbmanager.ConnectToDb()
	if dbErr != nil {
		log.Fatalln("(FATAL) >> Failed to connect to db: ", dbErr)
	}

	client, redisErr := redismanager.CreateClient()
	if redisErr != nil {
		log.Fatalln("(FATAL) >> Failed to create Redis client: ", redisErr)
	}

	repo := storage.NewPostgresRepo(dbPool)
	service := service.NewService(repo, client)
	handler := handler.NewHandler(&service)
	mux := http.NewServeMux()

	//ctx := context.Background()

	// urlString := "https://open.spotify.com/playlist/37i9dQZF1E8GuR0rPZpAQm"

	// shortCode, shortCodeErr := service.CreateShortCode(ctx, urlString)
	// if shortCodeErr != nil {
	// 	log.Println("(ERR) >> Failed to create short code: ", shortCodeErr)
	// 	return
	// }

	// _, getUrlErr := service.GetLongUrl(ctx, shortCode)
	// if getUrlErr != nil {
	// 	log.Println("(ERR) >> Failed to get long url: ", getUrlErr)
	// 	return
	// }

	mux.HandleFunc("POST /api/shorten", handler.ShortUrl)
	mux.HandleFunc("GET /{shortCode}", handler.Redirect)

	log.Println("Server is running on: http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln("(FATAL) >> Failed to run server: ", err)
	}
}
