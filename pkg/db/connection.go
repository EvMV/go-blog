package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewPostgresConnection() *gorm.DB {
	pgDsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.Open(pgDsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
