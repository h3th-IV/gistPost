package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Retriever interface {
	ConsumeAPI() ([]byte, error)
	ParseValues() (any, error)
	WriteToDB() error
}

func (p *Post) ConsumeAPI() ([]byte, error) {
	raw, err := MakeRequest(p.API)
	if err != nil {
		return nil, err
	}
	return raw, nil
}

func (p *Post) ParseValues() error {
	raw, err := p.ConsumeAPI()
	if err != nil {
		return err
	}
	// Unmarshal into the correct struct (p instead of NewPost)
	err = json.Unmarshal(raw, &p)
	if err != nil {
		return err
	}
	return nil
}
func (p *Post) WriteToDB() error {
	err := p.ParseValues()
	if err != nil {
		return err
	}
	query := `INSERT INTO posts(userid, id, title, body) values ($1, $2, $3, $4)`
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.UserId, p.Id, p.Title, p.Body)
	if err != nil {
		return err
	}
	return nil
}

func (c *Comments) ConsumeAPI() ([]byte, error) {
	raw, err := MakeRequest(c.API)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return raw, nil
}

func (c *Comments) ParseValues() error {
	raw, err := c.ConsumeAPI()
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, &c.comments)
	if err != nil {
		return err
	}
	for _, comment := range c.comments {
		fmt.Printf("PostID: %d\nID: %d\nName: %s\nEmail: %s\nBody: %s\n\n",
			comment.PostID, comment.ID, comment.Name, comment.Email, comment.Body)
		if comment.PostID == 2 {
			break
		}
	}
	return nil
}

func (c *Comments) WriteToDB() error {
	query := `insert into comments (postid, id, name, email, body) values ($1, $2, $3, $4, $5)`
	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	for _, comment := range c.comments {
		_, err = stmt.Exec(comment.PostID, comment.ID, comment.Name, comment.Email, comment.Body)
		if err != nil {
			return err
		}
	}
	return nil
}

// MakeRequest function
func MakeRequest(url string) ([]byte, error) {
	//create a new client
	client := &http.Client{}

	//create a new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	//send request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//read response body
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}
