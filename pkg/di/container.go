package di

import (
	"awesomeProject/internal/api/auth"
	"awesomeProject/pkg/route"
	"awesomeProject/pkg/server"
	"go.uber.org/fx"
	"log"
)

func BuildContainer() *fx.App {
	app := fx.New(
		fx.Provide(func() *log.Logger {
			return log.Default()
		}),
		fx.Provide(
			AsRoute(auth.NewLoginHandler),
			AsRoute(auth.NewRegisterHandler),
		),
		fx.Provide(
			fx.Annotate(
				server.NewAPIServer,
				fx.ParamTags(`group:"routes"`),
			),
		),
		fx.Invoke(func(apiServer *server.APIServer) {}),
	)

	return app
}

func AsRoute(handler any) any {
	return fx.Annotate(
		handler,
		fx.As(new(route.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
