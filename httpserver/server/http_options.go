package server

import (
	"net/http"
)

const defaultAddress = "127.0.0.1:18888"

type HttpOptions struct {
	Address string

	router http.Handler
}

type HttpOption func(*HttpOptions)

func newOptions(opts ...HttpOption) HttpOptions {
	opt := HttpOptions{}

	for _, o := range opts {
		o(&opt)
	}

	if len(opt.Address) == 0 {
		opt.Address = defaultAddress
	}

	return opt
}

func WithAddress(address string) HttpOption {
	return func(o *HttpOptions) {
		o.Address = address
	}
}

func WithRouter(r http.Handler) HttpOption {
	return func(o *HttpOptions) {
		o.router = r
	}
}
