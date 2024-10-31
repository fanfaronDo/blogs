package service

import (
	"github.com/fanfaronDo/blogs/internal/domain"
	"github.com/fanfaronDo/blogs/internal/repository"
)

type PostsService struct {
	repo *repository.Repository
}

func NewPostsService(repo *repository.Repository) *PostsService {
	return &PostsService{repo: repo}
}

func (p *PostsService) GetPosts(limit, offset int) ([]domain.Post, error) {
	return p.repo.Posts.GetPosts(limit, offset)
}

func (p *PostsService) GetAll() ([]domain.Post, error) {
	return p.repo.GetAll()
}

func (p *PostsService) GetTotal() (int, error) {
	return p.repo.GetTotal()
}
