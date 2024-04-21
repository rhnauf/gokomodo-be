package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	port := os.Getenv("DB_PORT")
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	dbConn, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("successfully connected to db on port =>", port)

	return &Database{Client: dbConn}, nil
}
