package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"./entity_db"

	"./repository_proj"
)

var (
	repo    repository_proj.PostRepository = repository_proj.NewPostRepository() //as our database store
	counter int                            = 2
)

func getPosts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "aplication/json")

	posts, err := repo.FindAll()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error" : "error getting the posts"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(posts)
}

func addPosts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "aplication/json")

	var post entity_db.Post

	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error" : "error unmarshaling requests}`))
		return
	}

	post.ID = int64(rand.Int())

	repo.Save(&post)
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(post)
}
