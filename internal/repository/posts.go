package repository

import (
	"database/sql"
	"github.com/fanfaronDo/blogs/internal/domain"
)

type PostsMysql struct {
	db *sql.DB
}

func NewPosts(db *sql.DB) *PostsMysql {
	return &PostsMysql{db: db}
}

func (p *PostsMysql) GetPosts(limit, offset int) ([]domain.Post, error) {
	var posts []domain.Post
	query := "SELECT title, content, image FROM posts LIMIT ? OFFSET ?;"
	rows, err := p.db.Query(query, limit, offset)
	if err != nil {
		return posts, err
	}

	defer rows.Close()
	for rows.Next() {
		var post domain.Post
		err = rows.Scan(&post.Title, &post.Content, &post.Image)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (p *PostsMysql) GetAll() ([]domain.Post, error) {
	var posts []domain.Post
	rows, err := p.db.Query("SELECT title, content, image FROM posts")
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post domain.Post
		err = rows.Scan(&post.Title, &post.Content, &post.Image)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return posts, err
	}

	return posts, nil
}

func (p *PostsMysql) GetTotal() (int, error) {
	var total int
	err := p.db.QueryRow("SELECT count(id) FROM posts;").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
