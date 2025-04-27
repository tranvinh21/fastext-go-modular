package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinhtran21/fastext-go-modular/infra/api/handler"
)

type Router struct {
	app            *fiber.App
	userHandler    *handler.UserHandler
	messageHandler *handler.MessageHandler
}

func NewRouter(app *fiber.App, userHandler *handler.UserHandler, messageHandler *handler.MessageHandler) *Router {
	return &Router{
		app:            app,
		userHandler:    userHandler,
		messageHandler: messageHandler,
	}
}

func (r *Router) RegisterRoutes() {
	// Health check endpoint
	r.app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// API v1 group
	v1 := r.app.Group("/api/v1")

	// User routes
	RegisterUserRoutes(v1, r.userHandler)
	RegisterMessageRoutes(v1, r.messageHandler)
}
