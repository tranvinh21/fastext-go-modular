package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinhtran21/fastext-go-modular/infra/api/handler"
)

func RegisterMessageRoutes(router fiber.Router, messageHandler *handler.MessageHandler) {
	messageGroup := router.Group("/messages")
	messageGroup.Get("/testing", messageHandler.HandleTesting)
}
