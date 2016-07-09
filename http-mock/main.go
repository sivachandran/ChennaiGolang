package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var (
	ApiUrl = "http://localhost:3000"
)

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func GetPosts() ([]Post, error) {
	resp, err := http.Get(ApiUrl + "/posts")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var posts []Post
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func NewPost(post *Post) error {
	var buf bytes.Buffer

	post.UserID = 101

	err := json.NewEncoder(&buf).Encode(post)
	if err != nil {
		return err
	}

	resp, err := http.Post(ApiUrl+"/posts", "application/json", &buf)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func main() {
	err := NewPost(
		&Post{
			ID:     rand.Intn(1000),
			UserID: rand.Intn(100),
			Title:  "Welcome Chennai Golang",
			Body:   fmt.Sprintf("Posted at %s", time.Now()),
		})
	if err != nil {
		panic(err)
	}

	p, err := GetPosts()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Posts: %#v\n", p)
}
