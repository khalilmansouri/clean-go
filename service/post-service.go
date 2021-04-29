package service

import (
	"clean-go/entity"
	"clean-go/repository"
	"errors"
	"math/rand"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	postRepository repository.PostRepository //= repository.NewFirestoreRepo()
)

func NewPostService(repository repository.PostRepository) PostService {
	postRepository = repository
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		return errors.New(":Post is empty")
	}
	if post.Text == "" {
		return errors.New(":Text is required")
	}
	if post.Title == "" {
		return errors.New(":Title post is empty")
	}
	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return postRepository.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return postRepository.FindAll()
}
