package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
	*Queries
}

func NewDB() *DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Printf("Connected to database")

	return &DB{
		DB:      db,
		Queries: New(db),
	}
}