package webServer

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vinhtran21/fastext-go-modular/config"
	"github.com/vinhtran21/fastext-go-modular/infra/api/handler"
	"github.com/vinhtran21/fastext-go-modular/infra/api/router"
	"go.uber.org/fx"
)

type HttpServer struct {
	app *fiber.App
}

func NewHttpServer(lc fx.Lifecycle, userHandler *handler.UserHandler, messageHandler *handler.MessageHandler) *HttpServer {
	app := fiber.New()
	router := router.NewRouter(app, userHandler, messageHandler)
	router.RegisterRoutes()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := app.Listen(":" + config.Envs.ServerConfig.Port); err != nil {
					// You can log here if necessary
					fmt.Printf("Error starting server: %v\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Shutting down HTTP server...")
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			fmt.Println("Waiting for 3 seconds for shutdown...")
			if err := app.ShutdownWithContext(ctx); err != nil {
				fmt.Println("Error during shutdown:", err)
				return err
			}

			fmt.Println("Shutdown completed within the timeout.")
			return nil
		},
	})

	return &HttpServer{
		app: app,
	}
}

func (h *HttpServer) StartServer() {
	h.app.Listen(":3000")
}
