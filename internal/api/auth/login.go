package auth

import (
	"fmt"
	"net/http"
)

type LoginHandler struct {
}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body []byte

	r.Body.Read(body)

	fmt.Println(string(body))
}

func (h *LoginHandler) Path() string {
	return "GET /login"
}
