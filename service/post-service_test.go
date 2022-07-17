package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yogie/go-clean-api/entity"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}
func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func (mock *MockRepository) FindByID(postID string) (*entity.Post, error) {
	// never been tested
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) Delete(post *entity.Post) error {
	// never been tested
	args := mock.Called()
	// result := args.Get(0)
	return args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {

	testService := NewPostService(nil)
	err := testService.Validate(nil)

	assert.NotNil(t, err)
}

func TestValidateEmptyTitle(t *testing.T) {
	post := entity.Post{
		ID:    1,
		Title: "",
		Text:  "Text",
	}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "the post title is empty", err.Error())
}

func TestFindAll(t *testing.T) {
	post := entity.Post{
		ID:    1,
		Title: "Title",
		Text:  "Text",
	}

	mockRepository := new(MockRepository)

	mockRepository.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepository)
	res, _ := testService.FindAll()

	mockRepository.AssertExpectations(t)

	assert.Equal(t, int64(1), res[0].ID)
	assert.Equal(t, "Title", res[0].Title)
	assert.Equal(t, "Text", res[0].Text)
}

func TestCreate(t *testing.T) {
	post := entity.Post{
		Title: "Title",
		Text:  "Text",
	}

	mockRepository := new(MockRepository)

	mockRepository.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepository)
	res, _ := testService.Create(&post)

	mockRepository.AssertExpectations(t)

	assert.NotNil(t, res.ID)
	assert.Equal(t, "Title", res.Title)
	assert.Equal(t, "Text", res.Text)
}
