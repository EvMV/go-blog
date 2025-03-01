package post

import (
	"awesomeProject/internal/application/interactor/post/createPost"
	"awesomeProject/internal/application/interactor/post/createPost/dto"
	"awesomeProject/internal/interfaceAdapters/http/provider"
	"encoding/json"
	"net/http"
)

type CreatePostHandler struct {
	interactor   *createPost.CreatePostInteractor
	userProvider provider.UserProvider
}

func NewCreatePostHandler(interactor *createPost.CreatePostInteractor, userProvider *provider.UserProvider) *CreatePostHandler {
	return &CreatePostHandler{interactor: interactor, userProvider: *userProvider}
}

func (h *CreatePostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := h.userProvider.ProvideCurrentUser(r)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var createPostDto dto.CreatePostDto

	if err := json.NewDecoder(r.Body).Decode(&createPostDto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	post := h.interactor.Execute(user, createPostDto)

	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}

func (h *CreatePostHandler) Path() string {
	return "POST /posts"
}
