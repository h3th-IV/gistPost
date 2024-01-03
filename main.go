package main

import (
	"fmt"
	"log"

	"github.com/h3th-IV/gistPost/database"
	"github.com/joho/godotenv"
)

func main() {
	//init env varaibles
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

	//get comments
	comments := database.Comments()
	for _, comment := range *comments {
		_, err = dB.Exec(`insert into comments (postid, id , name, email, body) values ($1, $2, $3, $4, $5)`, comment.PostID, comment.ID, comment.Name, comment.Email, comment.Body)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("PostID: %d\nID: %d\nName: %s\nEmail: %s\nBody: %s\n\n",
			comment.PostID, comment.ID, comment.Name, comment.Email, comment.Body)
		if comment.PostID == 2 {
			break
		}
	}

	//get user1 post
	userPost := database.Post()
	_, err = dB.Exec(`INSERT INTO posts(userid, id, title, body) values ($1, $2, $3, $4);`, userPost.UserId, userPost.Id, userPost.Title, userPost.Body)
	fmt.Println("postTItle: ", userPost.Title)
	if err != nil {
		log.Fatal(err)
	}
}
