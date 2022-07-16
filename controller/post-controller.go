package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/yogie/go-clean-api/entity"
	"github.com/yogie/go-clean-api/errors"
	"github.com/yogie/go-clean-api/service"
)

var (
	postService service.PostService
)

type postControllerStruct struct{}

type PostController interface {
	GetPosts(rw http.ResponseWriter, r *http.Request)
	AddPosts(rw http.ResponseWriter, r *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &postControllerStruct{}
}

func (*postControllerStruct) GetPosts(rw http.ResponseWriter, r *http.Request) {
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

func (*postControllerStruct) AddPosts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "aplication/json")

	var post entity.Post

	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(errors.ServiceError{Message: "error unmarshaling request"})

		return
	}

	post.ID = int64(rand.Intn(100000))

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
