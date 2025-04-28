package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	entity "github.com/vinhtran21/fastext-go-modular/domains/entities"
	usecase "github.com/vinhtran21/fastext-go-modular/domains/usecase"
	"github.com/vinhtran21/fastext-go-modular/infra/api/dto"
	"github.com/vinhtran21/fastext-go-modular/internal/util"
)

type AuthHandler struct {
	authUsecase *usecase.AuthUsecase
}

func NewAuthHandler(authUsecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	body := c.Locals("body").(*dto.LoginRequest)
	fmt.Println(body)
	user, err := h.authUsecase.Login(body.Email, body.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	accessToken, err := util.GenerateAccessToken(user.ID)
	refreshToken, err := util.GenerateRefreshToken(user.ID)

	util.SetCookie(c, "refreshToken", refreshToken, 30*24*60*60*1000)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Login successfula",
		"accessToken": accessToken,
	})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	body := c.Locals("body").(*dto.RegisterRequest)

	user := &entity.User{
		Email:    body.Email,
		Password: body.Password,
		Name:     body.Name,
	}

	err := h.authUsecase.Register(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Register successful",
	})
}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refreshToken")
	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Refresh token not found",
		})
	}

	userId, err := util.VerifyRefreshToken(refreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	accessToken, err := util.GenerateAccessToken(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Refresh token successful",
		"accessToken": accessToken,
	})
}
