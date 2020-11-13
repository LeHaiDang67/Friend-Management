package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

//InitDatabase is....
func InitDatabase() *sql.DB {
	db, err := getDatabase()
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(1)
	return db
}

func getDatabase() (*sql.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("Cannot get DATABASE_URL")
	}
	return sql.Open("postgres", databaseURL)
}
