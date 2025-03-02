package di

import (
	"fmt"
	"go.uber.org/fx"
	"goBlog/internal/application/interactor/auth/login"
	"goBlog/internal/application/interactor/auth/register"
	"goBlog/internal/application/interactor/post/createPost"
	"goBlog/internal/application/interactor/post/deletePost"
	"goBlog/internal/application/interactor/post/getAuthorPostList"
	iRepository "goBlog/internal/application/repository"
	"goBlog/internal/infrastructure/repository"
	"goBlog/internal/interfaceAdapters/http/handlers/auth"
	"goBlog/internal/interfaceAdapters/http/handlers/post"
	"goBlog/internal/interfaceAdapters/http/provider"
	"goBlog/pkg/db"
	"goBlog/pkg/route"
	"goBlog/pkg/server"
)

func ConfigureApp() *fx.App {
	app := fx.New(
		fx.Provide(db.NewPostgresConnection),
		fx.Provide(register.NewRegisterInteractor),
		fx.Provide(login.NewLoginInteractor),
		fx.Provide(createPost.NewCreatePostInteractor),
		fx.Provide(getAuthorPostList.NewGetAuthorPostListInteractor),
		fx.Provide(deletePost.NewDeletePostInteractor),
		fx.Provide(
			AsRequestHandler(auth.NewRegisterHandler),
			AsRequestHandler(auth.NewLoginHandler),
			AsRequestHandler(post.NewCreatePostHandler),
			AsRequestHandler(post.NewGetAuthorPostListHandler),
			AsRequestHandler(post.NewDeletePostHandler),
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
				fx.As(new(iRepository.UserRepository)),
			),
		),
		fx.Provide(provider.NewUserProvider),
		fx.Provide(provider.NewAuthorizeChecker),
		fx.Provide(
			fx.Annotate(
				repository.NewPgPostRepository,
				fx.As(new(iRepository.PostRepository)),
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
	err := s.Start()

	fmt.Println(err)
}
