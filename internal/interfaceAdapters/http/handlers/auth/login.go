package auth

import (
	"encoding/json"
	"goBlog/internal/application/interactor/auth/login"
	"goBlog/internal/application/interactor/auth/login/dto"
	"net/http"
)

type LoginHandler struct {
	interactor *login.LoginInteractor
}

func NewLoginHandler(interactor *login.LoginInteractor) *LoginHandler {
	return &LoginHandler{interactor}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var loginDto dto.LoginDto

	if err := json.NewDecoder(r.Body).Decode(&loginDto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	token, err := h.interactor.Execute(loginDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(token); err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}

func (h *LoginHandler) Path() string {
	return "POST /login"
}
