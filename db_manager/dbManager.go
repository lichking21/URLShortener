package dbmanager

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func GetConnection() (*pgx.Conn, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatalln("(FATAL) >> .env file wasn't found")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatalln("(FATAL) >> dbUrl is empty")
	}

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}

	return conn, err
}

func ConnectToDb() *pgx.Conn {

	conn, err := GetConnection()
	if err != nil {
		log.Fatalf("(FATAL) >> Не удалось установить соединение с БД: %v", err)
		return nil
	}

	defer conn.Close(context.Background())

	pingCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := conn.Ping(pingCtx); err != nil {
		log.Fatalf("(FATAL) >> База данных не отвечает на Ping: %v", err)
		return nil
	}

	log.Println("(LOG) >> Successfully connected to db")
	return conn
}
