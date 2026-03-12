package main

import (
	dbmanager "urlShortener/db_manager"
	redismanager "urlShortener/redis_manager"
)

func main() {

	dbmanager.ConnectToDb()
	redismanager.CreateClient()
}
