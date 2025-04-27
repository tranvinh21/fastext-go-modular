package repository

import (
	entity "github.com/vinhtran21/fastext-go-modular/domains/entities"
	"github.com/vinhtran21/fastext-go-modular/domains/repository"
	"github.com/vinhtran21/fastext-go-modular/infra/db"
)

type UserRepository struct {
	db *db.PostgresDB
}

func NewUserRepository(db *db.PostgresDB) repository.UserRepositoryImpl {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entity.User) error {
	return r.db.DB.Create(user).Error
}

func (r *UserRepository) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *entity.User) error {
	return r.db.DB.Save(user).Error
}

func (r *UserRepository) Delete(user *entity.User) error {
	return r.db.DB.Delete(user).Error
}
