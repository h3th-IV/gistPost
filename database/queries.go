package database

import (
	"encoding/json"
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

func Post() *post {
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
	var userPost post
	err = json.Unmarshal(raw, &userPost)
	if err != nil {
		log.Printf("Error parsing json: '%v'", err)
	}
	return &userPost

}

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func Comments() *[]Comment {
	//create new request
	request, err = http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/comments?postid=1", nil)
	if err != nil {
		log.Printf("Error making request: %v", err)
	}

	response, err = client.Do(request)
	if err != nil {
		log.Printf("Error Fetching response: %v", err)
	}

	defer response.Body.Close()

	raw, err = io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Err Reading Response: %v", err)
	}
	var comments []Comment
	err = json.Unmarshal(raw, &comments)
	if err != nil {
		log.Printf("Error decoding json ogject: %v", err)
	}
	return &comments
}
