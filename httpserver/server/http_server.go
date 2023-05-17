package server

import (
	"context"
	"net/http"
)

type HttpServer struct {
	opts HttpOptions

	srv http.Server
}

func NewHttpServer(opts ...HttpOption) Server {
	opt := newOptions(opts...)

	return &HttpServer{
		opts: opt,
	}
}

func (s *HttpServer) Start() error {
	srv := http.Server{
		Addr:    s.opts.Address,
		Handler: s.opts.router,
	}

	return srv.ListenAndServe()
}

func (s *HttpServer) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
