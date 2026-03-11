package main

import (
	"context"
	"log"
	"time"
	dbmanager "urlShortener/db_manager"
)

func main() {

	conn, err := dbmanager.ConnectToDb()
	if err != nil {
		log.Fatalf("(FATAL) >> Не удалось установить соединение с БД: %v", err)
	}

	defer conn.Close(context.Background())

	pingCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := conn.Ping(pingCtx); err != nil {
		log.Fatalf("(FATAL) >> База данных не отвечает на Ping: %v", err)
	}

	log.Println("(LOG) >> Successfully connected to db")
}
