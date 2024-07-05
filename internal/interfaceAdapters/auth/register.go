package auth

import (
	"awesomeProject/internal/application/interactor/auth/register"
	"awesomeProject/internal/application/interactor/auth/register/dto"
	"encoding/json"
	"net/http"
)

type RegisterHandler struct { // FIXME: handler or controller?
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

	user := h.interactor.Execute(registerDto)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}

func (h *RegisterHandler) Path() string {
	return "POST /register"
}
