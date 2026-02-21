package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn

func Init() error {
	//load env
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf(".envの読み込みに失敗: %w", err)
	}

	dsn := fmt.Sprintf(
		"host=localhost port=7777 user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return fmt.Errorf("DB接続に失敗: %w", err)
	}

	DB = conn
	return nil
}
