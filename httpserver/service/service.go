package service

import (
	"context"
	"time"

	"github.com/golang/glog"
)

type Service struct {
	opts Options
}

func NewService(opts ...Option) *Service {
	opt := newOptions(opts...)

	return &Service{
		opts: opt,
	}
}

func (s *Service) Start() error {
	go func() {
		if err := s.opts.server.Start(); err != nil {
			glog.Errorf("Server start error: %v", err)
		}
	}()

	return nil
}

func (s *Service) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.opts.server.Stop(ctx); err != nil {
		return err
	}

	return nil
}
