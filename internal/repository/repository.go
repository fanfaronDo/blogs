package repository

import (
	"database/sql"
	"github.com/fanfaronDo/blogs/internal/domain"
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

type Repository struct {
	Post
	Posts
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Post:  NewPost(db),
		Posts: NewPosts(db),
	}
}
