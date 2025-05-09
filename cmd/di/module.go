package di

import (
	usecase "github.com/vinhtran21/fastext-go-modular/domains/usecase"
	"github.com/vinhtran21/fastext-go-modular/infra/api/handler"
	"github.com/vinhtran21/fastext-go-modular/infra/api/router"
	"github.com/vinhtran21/fastext-go-modular/infra/db"
	"github.com/vinhtran21/fastext-go-modular/infra/repository"
	webServer "github.com/vinhtran21/fastext-go-modular/infra/web-server"
	"go.uber.org/fx"
)

var Repositories = fx.Options(
	fx.Provide(
		repository.NewUserRepository,
	),
)

var Usecases = fx.Options(
	fx.Provide(
		usecase.NewUserUsecase,
		usecase.NewMessageUsecase,
		usecase.NewAuthUsecase,
	),
)

var Handlers = fx.Options(
	fx.Provide(
		handler.NewUserHandler,
		handler.NewMessageHandler,
		handler.NewAuthHandler,
	),
)

var Module = fx.Options(
	fx.Provide(
		db.NewPostgresDBwFX,
		router.NewRouter,
		webServer.NewHttpServer,
	),
	Repositories,
	Usecases,
	Handlers,
)
