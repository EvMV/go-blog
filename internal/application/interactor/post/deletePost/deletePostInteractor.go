package deletePost

import (
	"awesomeProject/internal/application/repository"
	"awesomeProject/internal/domain/user"
	"errors"
)

type DeletePostInteractor struct {
	postRepository repository.PostRepository
}

func NewDeletePostInteractor(postRepository repository.PostRepository) *DeletePostInteractor {
	return &DeletePostInteractor{postRepository}
}

func (i DeletePostInteractor) Execute(user *user.User, postId int64) error {
	post := i.postRepository.FindOneById(postId)

	if post == nil {
		return errors.New("Post not found")
	}

	if post.AuthorID != user.Id {
		return errors.New("User not author")
	}

	i.postRepository.DeleteById(post.Id)

	return nil
}
