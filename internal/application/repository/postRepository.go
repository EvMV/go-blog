package repository

import "goBlog/internal/domain/post"

type PostRepository interface {
	Save(post *post.Post)
	FindAllByAuthorId(authorId int64) []*post.Post
	FindOneById(id int64) *post.Post
	DeleteById(id int64)
}
