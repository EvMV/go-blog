package repository

import "awesomeProject/internal/domain/post"

type PostRepository interface {
	Save(post *post.Post)
	FindAllByAuthorId(authorId int64) []*post.Post
}
