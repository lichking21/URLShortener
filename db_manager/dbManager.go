package dbmanager

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func ConnectToDb() (*pgx.Conn, error) {

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
