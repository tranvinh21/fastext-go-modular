package main

import (
	"github.com/vinhtran21/fastext-go-modular/cmd/di"
	webServer "github.com/vinhtran21/fastext-go-modular/infra/web-server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.Module,
		fx.Invoke(func(
			httpServer *webServer.HttpServer,
		) {
			httpServer.StartServer()
		}),
	).Run()
}
