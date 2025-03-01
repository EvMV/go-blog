package repository

import (
	"awesomeProject/internal/domain/post"
	"gorm.io/gorm"
)

type pgPostRepository struct {
	db *gorm.DB
}

func NewPgPostRepository(db *gorm.DB) *pgPostRepository {
	return &pgPostRepository{db}
}

func (r *pgPostRepository) Save(post *post.Post) {
	r.db.Save(post)
}

func (r *pgPostRepository) FindAllByAuthorId(authorId int64) []*post.Post {
	var posts []*post.Post
	r.db.Where("author_id = ?", authorId).Find(&posts)

	return posts
}
