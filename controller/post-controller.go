package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"../entity"
	"../errors"
	"../service"
)

var (
	postService service.PostService
)

type controller struct{}

type PostController interface {
	GetPosts(rw http.ResponseWriter, r *http.Request)
	AddPosts(rw http.ResponseWriter, r *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "aplication/json")

	posts, err := postService.FindAll()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(errors.ServiceError{Message: "error getting the posts"})

		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(posts)
}

func (*controller) AddPosts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "aplication/json")

	var post entity.Post

	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(errors.ServiceError{Message: "error unmarshaling request"})

		return
	}

	post.ID = rand.Int63()

	isNotValid := postService.Validate(&post)
	if isNotValid != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(errors.ServiceError{Message: isNotValid.Error()})
		return

	}
	result, err2 := postService.Create(&post)
	if err2 != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(errors.ServiceError{Message: "Error saving the post"})
		return

	}
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(result)
}
