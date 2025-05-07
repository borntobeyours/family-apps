package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	connStr := os.Getenv("DATABASE_URL")

	var err error
	DB, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	fmt.Println("Connected to DB with pool")
}
