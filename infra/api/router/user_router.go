package router

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/vinhtran21/fastext-go-modular/infra/api/handler"
)

func RegisterUserRoutes(router fiber.Router, userHandler *handler.UserHandler) {
	userGroup := router.Group("/users")
	userGroup.Get("/:email", userHandler.HandleGetUserByEmail)
	userGroup.Get("/testing", userHandler.HandleTesting)
}
