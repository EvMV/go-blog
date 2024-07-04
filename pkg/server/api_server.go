package server

import (
	"awesomeProject/pkg/route"
	"context"
	"go.uber.org/fx"
	"log"
	"net/http"
)

type APIServer struct {
	logger     *log.Logger
	httpServer *http.Server
}

func NewAPIServer(routes []route.Route, logger *log.Logger, lc fx.Lifecycle) *APIServer {
	mux := http.NewServeMux()

	for _, route := range routes {
		mux.Handle(route.Path(), route)
	}

	apiServer := &APIServer{
		logger: logger,
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go apiServer.Start()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return apiServer.httpServer.Shutdown(ctx)
		},
	})

	return apiServer
}

func (s *APIServer) Start() error {
	s.logger.Println("Starting server...")

	return s.httpServer.ListenAndServe()
}
