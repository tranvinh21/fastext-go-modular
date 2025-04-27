package webServer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinhtran21/fastext-go-modular/infra/api/handler"
	"github.com/vinhtran21/fastext-go-modular/infra/api/router"
)

type HttpServer struct {
	app *fiber.App
}

func NewHttpServer(userHandler *handler.UserHandler, messageHandler *handler.MessageHandler) *HttpServer {
	app := fiber.New()
	router := router.NewRouter(app, userHandler, messageHandler)
	router.RegisterRoutes()

	return &HttpServer{
		app: app,
	}
}

func (h *HttpServer) StartServer() {
	h.app.Listen(":3000")
}
