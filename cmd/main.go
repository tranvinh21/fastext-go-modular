package main

import (
	"github.com/vinhtran21/fastext-go-modular/cmd/di"
	webServer "github.com/vinhtran21/fastext-go-modular/infra/web-server"
	"go.uber.org/fx"
)

// close DB connection when program is closed

func main() {
	fx.New(
		di.Module,
		fx.Invoke(func(
			httpServer *webServer.HttpServer,
		) {
		},
		),
	).Run()
}
