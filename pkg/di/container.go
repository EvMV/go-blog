package di

import (
	"awesomeProject/internal/application/interactor/auth/register"
	repository2 "awesomeProject/internal/application/repository"
	"awesomeProject/internal/infrastructure/repository"
	"awesomeProject/internal/interfaceAdapters/auth"
	"awesomeProject/pkg/db"
	"awesomeProject/pkg/route"
	"awesomeProject/pkg/server"
	"go.uber.org/fx"
)

func ConfigureApp() *fx.App {
	app := fx.New(
		fx.Provide(db.NewPostgresConnection),
		fx.Provide(register.NewRegisterInteractor),
		fx.Provide(
			AsRequestHandler(auth.NewRegisterHandler),
		),
		fx.Provide(
			fx.Annotate(
				server.NewAPIServer,
				fx.ParamTags(`group:"routes"`),
			),
		),
		fx.Provide(
			fx.Annotate(
				repository.NewPgUserRepository,
				fx.As(new(repository2.UserRepository)),
			),
		),
		fx.Invoke(StartServer),
	)

	return app
}

func AsRequestHandler(handler any) any {
	return fx.Annotate(
		handler,
		fx.As(new(route.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

func StartServer(s *server.APIServer) {
	s.Start()
}
