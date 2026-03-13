package dbmanager

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectToDb() (*pgxpool.Pool, error) {

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return nil, errors.New("db url is empty")
	}

	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(pingCtx); err != nil {
		return nil, err
	}

	log.Println("(LOG) >> Successfully connected to db")
	return pool, nil
}
