package repository

import (
	"database/sql"
	"fmt"
	"github.com/fanfaronDo/blogs/internal/domain"
	"log"
)

type PostMysql struct {
	db *sql.DB
}

func NewPost(db *sql.DB) *PostMysql {
	return &PostMysql{
		db: db,
	}
}

func (p *PostMysql) Create(post domain.Post) error {
	query := "insert into posts (title, content, url, image) values (?, ?, ?, ?)"
	insert, err := p.db.Query(query, post.Title, post.Content, post.Url, post.Image)
	defer insert.Close()
	if err != nil {
		log.Printf("Post %s is not added: %v\n", post.Title, err)
		return err
	}

	return nil
}

func (p *PostMysql) Delete(post_id int) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	del, err := tx.Query("delete from posts where id=?", post_id)
	set, err := tx.Query("SET @newid := 0")
	update, err := tx.Query("UPDATE posts SET id = (@newid := @newid + 1)")

	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()

	if err != nil {
		return err
	}
	defer func() {
		del.Close()
		set.Close()
		update.Close()
	}()

	return nil
}

func (p PostMysql) GetById(post_id int) (domain.Post, error) {
	var prj domain.Post
	query := "SELECT title, content, image FROM posts WHERE id = ?"
	err := p.db.QueryRow(query, post_id).Scan(&prj.Title, &prj.Content, &prj.Image)

	if err != nil {
		if err == sql.ErrNoRows {
			return prj, fmt.Errorf("post with id %d not found", post_id)
		}
		return prj, err
	}

	return prj, nil
}

func (p PostMysql) Update(post_id int, post domain.Post) error {
	query := "update posts set title=$1, content=$1, image=$1 where id=?"
	update, err := p.db.Query(query, post.Title, post.Content, post.Image, post_id)
	if err != nil {
		log.Printf("Update post with id %d failed: %v\n", post, err)
		return err
	}
	defer update.Close()

	return nil
}
