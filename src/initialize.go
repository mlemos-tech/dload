package main

import (
	"mikaellemos.com.br/dload/src/config"
	"mikaellemos.com.br/dload/src/server"
)

type application struct {
	Server *server.Server
}

func InitializeApplication(propertie config.Config) (application, error) {
	handler := provideRouter()
	server := provideServer(handler, propertie)
	mainApplication := application{Server: server}
	return mainApplication, nil
}
