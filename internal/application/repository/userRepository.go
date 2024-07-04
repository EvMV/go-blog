package repository

import "awesomeProject/internal/domain"

type UserRepository interface {
	FindOneByEmail(email string) domain.User
	Save(user *domain.User)
}
