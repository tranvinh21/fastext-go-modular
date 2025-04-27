package usecase

import (
	entity "github.com/vinhtran21/fastext-go-modular/domains/entities"
	repository "github.com/vinhtran21/fastext-go-modular/domains/repository"
)

type UserUsecase struct {
	userRepository repository.UserRepositoryImpl
}

func NewUserUsecase(userRepository repository.UserRepositoryImpl) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (u *UserUsecase) FindByEmail(email string) (*entity.User, error) {
	return u.userRepository.FindByEmail(email)
}
