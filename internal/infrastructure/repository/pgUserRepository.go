package repository

import (
	"goBlog/internal/domain/user"
	"gorm.io/gorm"
)

type PgUserRepository struct {
	db *gorm.DB
}

func NewPgUserRepository(db *gorm.DB) *PgUserRepository {
	return &PgUserRepository{db}
}

func (r *PgUserRepository) FindOneByEmail(email string) *user.User {
	var user *user.User

	r.db.First(&user, "email = ?", email)

	return user
}

func (r *PgUserRepository) FindOneById(id float64) *user.User {
	var user *user.User

	r.db.First(&user, "id = ?", id)

	return user
}

func (r *PgUserRepository) Save(user *user.User) {
	r.db.Save(user)
}
