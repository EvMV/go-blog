package repository

import (
	"awesomeProject/internal/domain"
	"gorm.io/gorm"
)

type PgUserRepository struct {
	db *gorm.DB
}

func NewPgUserRepository(db *gorm.DB) *PgUserRepository {
	return &PgUserRepository{db}
}

func (r *PgUserRepository) FindOneByEmail(email string) domain.User {
	var user domain.User

	r.db.First(&user, email)

	return user
}

func (r *PgUserRepository) Save(user *domain.User) {
	r.db.Save(user)
}
