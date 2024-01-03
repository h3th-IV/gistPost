package main

import (
	"log"

	"github.com/h3th-IV/gistPost/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading files: %v", err)
	}
	// database.Post()
	database.Comments()
}
