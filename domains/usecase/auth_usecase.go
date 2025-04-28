package usecase

import (
	"errors"

	entity "github.com/vinhtran21/fastext-go-modular/domains/entities"
	"github.com/vinhtran21/fastext-go-modular/domains/repository"
	"github.com/vinhtran21/fastext-go-modular/internal/util"
	"gorm.io/gorm"
)

type AuthUsecase struct {
	userRepository repository.UserRepositoryImpl
}

func NewAuthUsecase(userRepository repository.UserRepositoryImpl) *AuthUsecase {
	return &AuthUsecase{userRepository: userRepository}
}

func (u *AuthUsecase) Login(email string, password string) (entity.User, error) {
	foundUser, err := u.userRepository.FindByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return entity.User{}, err
	}

	if foundUser == nil {
		password = "dummy_invalid_hash_for_timing_attack_mitigation"
	} else {
		password = foundUser.Password
	}

	if !util.VerifyPassword(password, foundUser.Password) {
		return entity.User{}, errors.New("invalid credentials")
	}

	return *foundUser, nil
}

func (u *AuthUsecase) Register(user *entity.User) error {
	foundUser, err := u.userRepository.FindByEmail(user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if foundUser != nil {
		return errors.New("email already exists")
	}

	foundUser, err = u.userRepository.FindByUsername(user.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if foundUser != nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return u.userRepository.Create(user)
}
