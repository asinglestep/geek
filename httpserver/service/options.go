package service

import "geek/httpserver/server"

type Options struct {
	server server.Server
}

func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

type Option func(*Options)

func WithServer(srv server.Server) Option {
	return func(o *Options) {
		o.server = srv
	}
}
