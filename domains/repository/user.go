package repository

import (
	entity "github.com/vinhtran21/fastext-go-modular/domains/entities"
)

type UserRepositoryImpl interface {
	Create(user *entity.User) error
	FindByID(id uint) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(user *entity.User) error
}
