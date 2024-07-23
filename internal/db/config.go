package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Config struct {
	ConnectionString string `conf:"connection_string"`
}

func Connect(cfg Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.ConnectionString)
	if err != nil {
		log.Fatalln("failed to open database connection:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("failed to ping database:", err)
	}

	return db
}
