package getAuthorPostList

import (
	"awesomeProject/internal/application/repository"
	"awesomeProject/internal/domain/post"
)

type GetAuthorPostList struct {
	postRepository repository.PostRepository
}

func NewGetAuthorPostList(postRepository repository.PostRepository) *GetAuthorPostList {
	return &GetAuthorPostList{postRepository: postRepository}
}

func (i *GetAuthorPostList) Execute(authorId int64) []*post.Post {
	posts := i.postRepository.FindAllByAuthorId(authorId)

	return posts
}
