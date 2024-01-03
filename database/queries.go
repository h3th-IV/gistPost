package database

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var (
	client   = &http.Client{}
	request  *http.Request
	response *http.Response
	raw      []byte
)

func Post() {
	//create new request
	request, err = http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		log.Fatalf("Error making request: '%v'\n", err)
	}
	//sedn request via client
	response, err = client.Do(request)
	if err != nil {
		log.Fatal("Err Getting response")
	}
	defer response.Body.Close()
	raw, err = io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Err getting Data: %v", err)
	}
	fmt.Println(string(raw))
	var userPost post
	err = json.Unmarshal(raw, &userPost)
	if err != nil {
		log.Printf("Error parsing json: '%v'", err)
	}
	fmt.Print(userPost.Title)
	//writing to dB...and stuff
	dB, err = InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer CloseDB()

	_, err = dB.Exec(`INSERT INTO posts(userid, id, title, body) values ($1, $2, $3, $4);`, userPost.UserId, userPost.Id, userPost.Title, userPost.Body)
	if err != nil {
		log.Fatal(err)
	}
}

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func Comments() {
	//create new request
	request, err = http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/comments?postid=1", nil)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	response, err = client.Do(request)
	if err != nil {
		log.Fatalf("Error Fetching response: %v", err)
	}

	defer response.Body.Close()

	raw, err = io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Err Reading Response: %v", err)
	}
	var comments []Comment
	err = json.Unmarshal(raw, &comments)
	if err != nil {
		log.Fatalf("Error decoding json ogject: %v", err)
	}

	dB, err = InitDB()
	if err != nil {
		log.Fatalf("Err connecting to database: %v", err)
	}
	defer CloseDB()

	for _, comment := range comments {
		_, err = dB.Exec(`insert into comments (postid, id , name, email, body) values ($1, $2, $3, $4, $5)`, comment.PostID, comment.ID, comment.Name, comment.Email, comment.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("PostID: %d\nID: %d\nName: %s\nEmail: %s\nBody: %s\n\n",
			comment.PostID, comment.ID, comment.Name, comment.Email, comment.Body)
		if comment.PostID == 2 {
			break
		}
	}
}
