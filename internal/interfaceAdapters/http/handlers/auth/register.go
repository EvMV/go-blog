package auth

import (
	"encoding/json"
	"goBlog/internal/application/interactor/auth/register"
	"goBlog/internal/application/interactor/auth/register/dto"
	"net/http"
)

type RegisterHandler struct {
	interactor *register.RegisterInteractor
}

func NewRegisterHandler(interactor *register.RegisterInteractor) *RegisterHandler {
	return &RegisterHandler{interactor}
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var registerDto dto.RegisterDto

	if err := json.NewDecoder(r.Body).Decode(&registerDto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.interactor.Execute(registerDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}

func (h *RegisterHandler) Path() string {
	return "POST /register"
}
