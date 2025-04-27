package handler

import (
	"github.com/gofiber/fiber/v2"
	usecase "github.com/vinhtran21/fastext-go-modular/domains/usecase"
)

type MessageHandler struct {
	messageUsecase *usecase.MessageUsecase
}

func NewMessageHandler(messageUsecase *usecase.MessageUsecase) *MessageHandler {
	return &MessageHandler{
		messageUsecase: messageUsecase,
	}
}

func (h *MessageHandler) HandleTesting(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}
