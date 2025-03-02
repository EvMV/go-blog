package getAuthorPostList

import (
	"awesomeProject/internal/application/repository"
	"awesomeProject/internal/domain/post"
)

type GetAuthorPostListInteractor struct {
	postRepository repository.PostRepository
}

func NewGetAuthorPostListInteractor(postRepository repository.PostRepository) *GetAuthorPostListInteractor {
	return &GetAuthorPostListInteractor{postRepository: postRepository}
}

func (i *GetAuthorPostListInteractor) Execute(authorId int64) []*post.Post {
	posts := i.postRepository.FindAllByAuthorId(authorId)

	return posts
}
