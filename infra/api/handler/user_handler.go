package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinhtran21/fastext-go-modular/domains/usecase"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) HandleGetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	user, err := h.userUsecase.FindByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleTesting(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}
