package createPost

import (
	"goBlog/internal/application/interactor/post/createPost/dto"
	"goBlog/internal/application/repository"
	"goBlog/internal/domain/post"
	"goBlog/internal/domain/user"
)

type CreatePostInteractor struct {
	postRepository repository.PostRepository
}

func NewCreatePostInteractor(postRepository repository.PostRepository) *CreatePostInteractor {
	return &CreatePostInteractor{postRepository: postRepository}
}

func (i *CreatePostInteractor) Execute(user *user.User, dto dto.CreatePostDto) *post.Post {
	post := post.NewPost(dto.Title, dto.Content, user)

	i.postRepository.Save(post)

	return post
}
