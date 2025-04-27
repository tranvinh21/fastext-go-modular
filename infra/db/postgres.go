package db

import (
	"context"
	"fmt"

	"github.com/vinhtran21/fastext-go-modular/config"
	"go.uber.org/fx"
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
func NewPostgresDBwFX(lc fx.Lifecycle) (*PostgresDB, error) {
	fmt.Println("NewPostgresDBwFX")
	db, err := gorm.Open(postgres.Open(config.Envs.DBConfig.URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting DB connection...")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Closing DB connection...")
			sqlDB, _ := db.DB()
			return sqlDB.Close()
		},
	})
	return &PostgresDB{DB: db}, nil
}
