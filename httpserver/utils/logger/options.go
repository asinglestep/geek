package logger

type Options struct {
	Level string
}

func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	if len(opt.Level) == 0 {
		opt.Level = "debug"
	}

	return opt
}

type Option func(*Options)

func WithLevel(l string) Option {
	return func(o *Options) {
		o.Level = l
	}
}
