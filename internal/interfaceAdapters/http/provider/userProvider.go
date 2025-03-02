package provider

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"goBlog/internal/application/repository"
	"goBlog/internal/domain/user"
	"net/http"
)

type UserProvider struct {
	userRepository repository.UserRepository
}

func NewUserProvider(userRepository repository.UserRepository) *UserProvider {
	return &UserProvider{userRepository: userRepository}
}

func (p *UserProvider) ProvideCurrentUser(request *http.Request) (*user.User, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(request.Header.Get("Authorization"), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("Ws5Jhf"), nil
	})

	if err != nil {
		return nil, err
	}

	userId := claims["sub"].(float64)

	user := p.userRepository.FindOneById(userId)

	if user == nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}
