package post

import (
	"awesomeProject/internal/application/interactor/post/getAuthorPostList"
	"awesomeProject/internal/interfaceAdapters/http/provider"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type GetAuthorPostListHandler struct {
	interactor *getAuthorPostList.GetAuthorPostListInteractor
	auhorizer  *provider.AuthorizeChecker
}

func NewGetAuthorPostListHandler(
	interactor *getAuthorPostList.GetAuthorPostListInteractor,
	auhorizer *provider.AuthorizeChecker,
) *GetAuthorPostListHandler {
	return &GetAuthorPostListHandler{interactor: interactor, auhorizer: auhorizer}
}

func (h *GetAuthorPostListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.auhorizer.Authorize(r); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	userId, err := h.getAuthorIdFromPath(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	posts := h.interactor.Execute(userId)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}

func (h *GetAuthorPostListHandler) Path() string {
	return "GET /posts/{authorId}"
}

func (h *GetAuthorPostListHandler) getAuthorIdFromPath(r *http.Request) (int64, error) {
	pathUserId := r.PathValue("authorId")
	if pathUserId == "" {
		return 0, fmt.Errorf("Invalid path value")
	}

	userId, err := strconv.ParseInt(pathUserId, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid path value")
	}

	return userId, nil
}
