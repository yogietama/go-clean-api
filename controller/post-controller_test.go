package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yogie/go-clean-api/entity"
	"github.com/yogie/go-clean-api/repository"
	"github.com/yogie/go-clean-api/service"
)

var (
	postRepoTest       repository.PostRepository = repository.NewPostgreRepository()
	postServiceTest    service.PostService       = service.NewPostService(postRepoTest)
	postControllerTest PostController            = NewPostController(postServiceTest)
)

const (
	ID    int64  = 123
	TITLE string = "Test Title"
	TEXT  string = "Test Text"
)

func TestAddPosts(t *testing.T) {
	// create a new HTTP POST request
	var jsonValue = []byte(`{"title" : "Test Title", "text" : "Test Text"}`)

	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonValue))

	// assign HTTP handler function (Add post function)
	handler := http.HandlerFunc(postControllerTest.AddPosts)

	// record HTTP response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// add assertion on http status code and the response
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: gor %v want %v", status, http.StatusOK)
	}

	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	// assertionn HTTP response
	assert.NotNil(t, post.ID)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	// cleanUp testing data in database
	postRepoTest.Delete(&post)

}

func TestGetPosts(t *testing.T) {
	// insert data test
	setup()

	// create a get HTTP request
	req, _ := http.NewRequest("GET", "/posts", nil)

	// assign HTTP handler function (Get post function)
	handler := http.HandlerFunc(postControllerTest.GetPosts)

	// record HTTP response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// add assertion on http status code and the response
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: gor %v want %v", status, http.StatusOK)
	}

	var posts []entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	len_posts := len(posts) - 1

	// assertionn HTTP response
	assert.NotNil(t, posts[len_posts].ID)
	assert.Equal(t, TITLE, posts[len_posts].Title)
	assert.Equal(t, TEXT, posts[len_posts].Text)

	// cleanUp testing data in database
	postRepoTest.Delete(&posts[len_posts])

}

func setup() {
	var post entity.Post = entity.Post{
		ID:    ID,
		Title: TITLE,
		Text:  TEXT,
	}

	postRepoTest.Save(&post)
}
