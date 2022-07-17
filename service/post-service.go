package service

import (
	"errors"

	"github.com/yogie/go-clean-api/entity"
	"github.com/yogie/go-clean-api/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByID(postID string) (*entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService(repoval repository.PostRepository) PostService {
	repo = repoval
	return &service{}
}

func (*service) FindByID(postID string) (*entity.Post, error) {
	return repo.FindByID(postID)
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("the post is empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("the post title is empty")
		return err
	}

	return nil

}
func (*service) Create(post *entity.Post) (*entity.Post, error) {
	return repo.Save(post)
}
func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
