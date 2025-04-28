package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/vinhtran21/fastext-go-modular/infra/api/dto"
	"github.com/vinhtran21/fastext-go-modular/infra/api/handler"
	"github.com/vinhtran21/fastext-go-modular/infra/api/middleware"
)

var validate = validator.New()

func RegisterAuthRoutes(router fiber.Router, authHandler *handler.AuthHandler) {
	authRouter := router.Group("/auth")
	authRouter.Post("/signin", middleware.ValidateBody[dto.LoginRequest](), authHandler.Login)
	authRouter.Post("/signup", middleware.ValidateBody[dto.RegisterRequest](), authHandler.Register)
	authRouter.Post("/refresh-token", authHandler.RefreshToken)
}
