package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"geek/httpserver/config"
	_ "geek/httpserver/docs"
	"geek/httpserver/router"
	"geek/httpserver/server"
	"geek/httpserver/service"
	"geek/httpserver/utils/logger"

	"github.com/BurntSushi/toml"
)

var configPath string

func main() {
	flag.StringVar(&configPath, "config", "./etc/config.toml", "config path")
	flag.Parse()

	var conf config.Config
	if _, err := toml.DecodeFile(configPath, &conf); err != nil {
		logger.Errorf("toml.DecodeFile error: %v", err)
		os.Exit(-1)
	}

	logger.Infof("config: %+v", conf)
	l := logger.NewLogger(logger.WithLevel(conf.Log.Level))
	logger.SetLogger(l)

	r := router.Setup()

	srv := server.NewHttpServer(server.WithAddress(conf.Server.Address), server.WithRouter(r))
	svc := service.NewService(service.WithServer(srv))

	if err := svc.Start(); err != nil {
		logger.Errorf("svc.Start error: %v", err)
		os.Exit(-1)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	s := <-quit
	logger.Infof("got signal %v", s)
	if err := svc.Stop(); err != nil {
		logger.Errorf("svc.Stop error: %v", err)
		os.Exit(-1)
	}
}
