package database

import (
	"context"
	"database/sql"
	"log"
	"os"
)

var (
	SQLITE_URL = "SQLITE_URL"
)

func NewSqliteConnection(ctx context.Context) (*sql.DB, error) {
	sqlite3Uri := os.Getenv(SQLITE_URL)
	client, err := sql.Open("sqlite3", sqlite3Uri)

	if err != nil {
		println("Deu pau no banco.")
		log.Fatal(err.Error())
	}

	if err := client.PingContext(ctx); err != nil {
		return nil, err
	}

	return client, nil
}
