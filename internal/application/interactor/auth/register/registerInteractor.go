package register

import (
	"awesomeProject/internal/application/interactor/auth/register/dto"
	"awesomeProject/internal/application/repository"
	"awesomeProject/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInteractor struct {
	userRepository repository.UserRepository
}

func NewRegisterInteractor(userRepository repository.UserRepository) *RegisterInteractor {
	return &RegisterInteractor{userRepository}
}

func (r RegisterInteractor) Execute(dto dto.RegisterDto) *domain.User {
	password, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err) // FIXME
	}

	user := domain.NewUser(dto.Name, dto.Email, string(password))
	r.userRepository.Save(user)

	return user
}
