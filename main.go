package main

import (
	"log"

	"github.com/h3th-IV/gistPost/database"
	"github.com/h3th-IV/gistPost/internal"
	"github.com/joho/godotenv"
)

func main() {
	// init env varaibles
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading files: %v", err)
	}

	//start database
	dB, err := database.InitDB()
	if err != nil {
		log.Printf("Err connecting to database: %v", err)
	}
	//always remember to shut it down
	defer database.CloseDB()
	var (
		// post = &internal.Post{
		// 	API: "https://jsonplaceholder.typicode.com/posts/1",
		// 	DB:  dB,
		// }

		comments = &internal.Comments{
			API: "https://jsonplaceholder.typicode.com/comments?postid=1",
			DB:  dB,
		}
	)
	// post.WriteToDB()
	comments.WriteToDB()
}
