package service

import (
	"github.com/fanfaronDo/blogs/internal/domain"
	"github.com/fanfaronDo/blogs/internal/repository"
)

type PostService struct {
	repo *repository.Repository
}

func NewPostService(repo *repository.Repository) *PostService {
	return &PostService{repo}
}

func (p *PostService) Create(post domain.Post) error {
	return p.repo.Create(post)
}

func (p *PostService) Update(post_id int, post domain.Post) error {
	return p.repo.Update(post_id, post)
}

func (p *PostService) GetById(post_id int) (domain.Post, error) {
	return p.repo.GetById(post_id)
}

func (p *PostService) Delete(post_id int) error {
	return p.repo.Delete(post_id)
}
