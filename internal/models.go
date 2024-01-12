package internal

import "database/sql"

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	API    string
	DB     *sql.DB
}

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

type Comments struct {
	API      string
	DB       *sql.DB
	comments []Comment
}
