package service

import (
	"github.com/fanfaronDo/blogs/internal/domain"
	"github.com/fanfaronDo/blogs/internal/repository"
)

type Posts interface {
	GetAll() ([]domain.Post, error)
	GetTotal() (int, error)
	GetPosts(limit, offset int) ([]domain.Post, error)
}

type Post interface {
	Create(post domain.Post) error
	Delete(post_id int) error
	GetById(post_id int) (domain.Post, error)
	Update(post_id int, post domain.Post) error
}

type Service struct {
	Post
	Posts
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Post:  NewPostService(repos),
		Posts: NewPostsService(repos),
	}
}
