package repository

import (
	"awesomeProject/internal/domain/post"
	"fmt"
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
	r.db.Where("author_id = ?", authorId).Preload("Author").Find(&posts)
	fmt.Println(posts[0])
	return posts
}

func (r *pgPostRepository) FindOneById(id int64) *post.Post {
	var post post.Post
	r.db.First(&post, id)

	return &post
}

func (r *pgPostRepository) DeleteById(id int64) {
	r.db.Delete(&post.Post{}, id)
}
