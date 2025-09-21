package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool // bu neden dışarıda

var dbURL = os.Getenv("DB_URL")

func ConnectDB() {
	var err error
	Pool, err = pgxpool.New(context.Background(),
		dbURL)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}
	log.Println("Connected to Postgres ✅")
}
