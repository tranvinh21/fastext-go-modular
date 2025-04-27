package usecase

import (
	"fmt"

	"github.com/vinhtran21/fastext-go-modular/domains/repository"
)

type MessageUsecase struct {
	messageRepository repository.UserRepositoryImpl
}

func NewMessageUsecase(messageRepository repository.UserRepositoryImpl) *MessageUsecase {
	return &MessageUsecase{
		messageRepository: messageRepository,
	}
}

func (u *MessageUsecase) Testing() {
	fmt.Println("Testing")
}
