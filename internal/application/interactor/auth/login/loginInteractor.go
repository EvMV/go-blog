package login

import (
	"awesomeProject/internal/application/interactor/auth/login/dto"
	"awesomeProject/internal/application/repository"
	"awesomeProject/internal/domain/user"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginInteractor struct {
	userRepository repository.UserRepository
}

func NewLoginInteractor(userRepository repository.UserRepository) *LoginInteractor {
	return &LoginInteractor{userRepository}
}

func (l *LoginInteractor) Execute(dto dto.LoginDto) (string, error) {
	user := l.userRepository.FindOneByEmail(dto.Email)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return "", errors.New("incorrect password")
	}

	jwtToken := createToken(user)

	return jwtToken.SignedString([]byte("Ws5Jhf"))
}

func createToken(user *user.User) *jwt.Token {
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	claims := jwtToken.Claims.(jwt.MapClaims)
	claims["sub"] = user.Id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return jwtToken
}
