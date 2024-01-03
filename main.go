package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading files: %v", err)
	}
	// database.Post()
	// database.Comments()
}
