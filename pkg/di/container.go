package di

import (
	"awesomeProject/internal/application/interactor/auth/login"
	"awesomeProject/internal/application/interactor/auth/register"
	iUserRepository "awesomeProject/internal/application/repository"
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
		fx.Provide(login.NewLoginInteractor),
		fx.Provide(
			AsRequestHandler(auth.NewRegisterHandler),
			AsRequestHandler(auth.NewLoginHandler),
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
				fx.As(new(iUserRepository.UserRepository)),
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
