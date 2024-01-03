package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	dB  *sql.DB
	err error
)

func InitDB() (*sql.DB, error) {
	connstring := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"), os.Getenv("PG_DBNAME"))

	dB, err = sql.Open("postgres", connstring)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return nil, err
	}

	// Set up connection pool

	err = dB.Ping()
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}

	log.Println("Database connection initialized successfully")
	return dB, nil
}

func CloseDB() {
	if dB != nil {
		dB.Close()
		log.Println("Database connection closed Succesfully")
	}
}
