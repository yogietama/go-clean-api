package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts   []Post // as our database store
	counter int    = 2
)

func init() {
	posts = []Post{{Id: 1, Title: "Title 1", Text: "Text 1"}}
}

func getPosts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "aplication/json")

	result, err := json.Marshal(posts)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error" : "error marshaling the posts array"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(result)
}

func addPosts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "aplication/json")

	var post Post

	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error" : "error unmarshaling requests}`))
		return
	}

	post.Id = counter
	counter++

	posts = append(posts, post)

	result, err := json.Marshal(post)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error" : "error marshaling the posts array"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(result)
}
