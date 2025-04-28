package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinhtran21/fastext-go-modular/infra/api/handler"
)

type Router struct {
	app            *fiber.App
	userHandler    *handler.UserHandler
	messageHandler *handler.MessageHandler
	authHandler    *handler.AuthHandler
}

func NewRouter(app *fiber.App, userHandler *handler.UserHandler, messageHandler *handler.MessageHandler, authHandler *handler.AuthHandler) *Router {
	return &Router{
		app:            app,
		userHandler:    userHandler,
		messageHandler: messageHandler,
		authHandler:    authHandler,
	}
}

func (r *Router) RegisterRoutes() {
	// Health check endpoint
	r.app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("sOK")
	})

	apiRoutes := r.app.Group("/api")

	// User routes
	RegisterUserRoutes(apiRoutes, r.userHandler)
	RegisterMessageRoutes(apiRoutes, r.messageHandler)
	RegisterAuthRoutes(apiRoutes, r.authHandler)
}
