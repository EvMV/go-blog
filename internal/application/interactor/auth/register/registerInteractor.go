package register

import (
	"goBlog/internal/application/interactor/auth/register/dto"
	"goBlog/internal/application/repository"
	"goBlog/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInteractor struct {
	userRepository repository.UserRepository
}

func NewRegisterInteractor(userRepository repository.UserRepository) *RegisterInteractor {
	return &RegisterInteractor{userRepository}
}

func (r RegisterInteractor) Execute(dto dto.RegisterDto) (*user.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := user.NewUser(dto.Name, dto.Email, string(password))
	r.userRepository.Save(user)

	return user, nil
}
