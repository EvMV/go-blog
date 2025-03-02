package server

import (
	"goBlog/pkg/route"
	"net/http"
)

type APIServer struct {
	httpServer *http.Server
}

func NewAPIServer(routes []route.Route) *APIServer {
	mux := http.NewServeMux()

	for _, route := range routes {
		mux.Handle(route.Path(), route)
	}

	apiServer := &APIServer{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
	}

	return apiServer
}

func (s *APIServer) Start() error {
	return s.httpServer.ListenAndServe()
}
