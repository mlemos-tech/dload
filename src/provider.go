package main

import (
	"github.com/fate-lovely/phi"
	"mikaellemos.com.br/dload/src/config"
	"mikaellemos.com.br/dload/src/server"
	"mikaellemos.com.br/dload/src/web"
)

func provideServer(handler phi.Handler, config config.Properties) *server.Server {
	return &server.Server{
		Host:    config.Server.Host,
		Port:    config.Server.Port,
		Handler: handler,
	}
}

func provideRouter() phi.Handler {
	r := phi.NewRouter()
	r.Mount("/", web.Handler())
	return r
}
