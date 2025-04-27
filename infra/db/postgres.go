package db

import (
	"github.com/vinhtran21/fastext-go-modular/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

func NewPostgresDB() (*PostgresDB, error) {
	db, err := gorm.Open(postgres.Open(config.Envs.DBConfig.URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PostgresDB{DB: db}, nil
}
