package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"

	"github.com/yogie/go-clean-api/cache"
	"github.com/yogie/go-clean-api/entity"
	"github.com/yogie/go-clean-api/errors"
	"github.com/yogie/go-clean-api/service"
)

var (
	postService service.PostService
	postCache   cache.PostCache
)

type postControllerStruct struct{}

type PostController interface {
	GetPostByID(rw http.ResponseWriter, r *http.Request)
	GetPosts(rw http.ResponseWriter, r *http.Request)
	AddPosts(rw http.ResponseWriter, r *http.Request)
}

func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	postService = service
	postCache = cache
	return &postControllerStruct{}
}
func (*postControllerStruct) GetPostByID(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "aplication/json")
	postID := strings.Split(r.URL.Path, "/")[2]

	var post *entity.Post = postCache.Get(postID)
	if post == nil {
		post, err := postService.FindByID(postID)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			json.NewEncoder(rw).Encode(errors.ServiceError{
				Message: "No posts found",
			})
		} else {
			postCache.Set(postID, post)
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(post)
		}
	} else {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(post)
	}
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
