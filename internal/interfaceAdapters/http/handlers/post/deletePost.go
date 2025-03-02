package post

import (
	"fmt"
	"goBlog/internal/application/interactor/post/deletePost"
	"goBlog/internal/interfaceAdapters/http/provider"
	"net/http"
	"strconv"
)

type DeletePostHandler struct {
	userProvider *provider.UserProvider
	interactor   *deletePost.DeletePostInteractor
}

func NewDeletePostHandler(userProvider *provider.UserProvider, interactor *deletePost.DeletePostInteractor) *DeletePostHandler {
	return &DeletePostHandler{userProvider, interactor}
}

func (h *DeletePostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := h.userProvider.ProvideCurrentUser(r)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	postId, err := h.getPostId(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.interactor.Execute(user, postId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *DeletePostHandler) Path() string {
	return "DELETE /posts/{postId}"
}

func (h *DeletePostHandler) getPostId(r *http.Request) (int64, error) {
	pathPostId := r.PathValue("postId")
	if pathPostId == "" {
		return 0, fmt.Errorf("Invalid path value")
	}

	postId, err := strconv.ParseInt(pathPostId, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid path value")
	}

	return postId, nil
}
