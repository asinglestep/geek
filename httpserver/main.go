package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"geek/httpserver/config"
	"geek/httpserver/router"
	"geek/httpserver/server"
	"geek/httpserver/service"

	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	var conf config.Config
	if _, err := toml.DecodeFile("./etc/config.toml", &conf); err != nil {
		glog.Errorf("toml.DecodeFile error: %v", err)
		os.Exit(-1)
	}

	glog.Infof("config: %+v", conf)
	r := router.Setup()

	srv := server.NewHttpServer(server.WithAddress(conf.Server.Address), server.WithRouter(r))
	svc := service.NewService(service.WithServer(srv))

	if err := svc.Start(); err != nil {
		glog.Errorf("svc.Start error: %v", err)
		os.Exit(-1)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	s := <-quit
	glog.Infof("got signal %v", s)

	if err := svc.Stop(); err != nil {
		glog.Errorf("svc.Stop error: %v", err)
		os.Exit(-1)
	}
}
