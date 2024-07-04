package auth

import (
	"fmt"
	"net/http"
)

type RegisterHandler struct {
}

func NewRegisterHandler() *RegisterHandler {
	return &RegisterHandler{}
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body []byte

	r.Body.Read(body)

	fmt.Println(string(body))
}

func (h *RegisterHandler) Path() string {
	return "GET /register"
}
