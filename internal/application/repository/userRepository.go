package repository

import (
	"goBlog/internal/domain/user"
)

type UserRepository interface {
	FindOneByEmail(email string) *user.User
	FindOneById(id float64) *user.User
	Save(user *user.User)
}
