package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetPosts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	posts := []Post{
		{ID: 1, UserID: 10, Title: "Test Title 1", Body: "Test Body 1"},
		{ID: 2, UserID: 20, Title: "Test Title 2", Body: "Test Body 2"},
	}

	httpmock.RegisterResponder("GET", "http://localhost:3000/posts",
		func(req *http.Request) (*http.Response, error) {
			/*resp, err := httpmock.NewStringResponse(200, posts, "{id:1, userId:10, }")
			if err != nil {
				return nil, err
			}*/
			resp, err := httpmock.NewJsonResponse(200, posts)
			if err != nil {
				return nil, err
			}
			return resp, nil
		})

	p, err := GetPosts()
	assert.NoError(t, err, "error while getting posts")
	assert.Equal(t, p, posts, "posts mismatch")
}

func TestNewPost(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	var p Post
	httpmock.RegisterResponder("POST", "http://localhost:3000/posts",
		func(req *http.Request) (*http.Response, error) {
			err := json.NewDecoder(req.Body).Decode(&p)
			if err != nil {
				return nil, err
			}

			return httpmock.NewStringResponse(201, ""), nil
		})

	post := Post{ID: 1, UserID: 10, Title: "Test Title 1", Body: "Test Body 1"}
	err := NewPost(&post)
	assert.NoError(t, err, "error while posting post")
	assert.Equal(t, p, post, "post mismatch")
}
