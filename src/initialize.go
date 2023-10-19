package main

import (
	"mikaellemos.com.br/dload/src/config"
	"mikaellemos.com.br/dload/src/migrate"
	"mikaellemos.com.br/dload/src/repository"
	"mikaellemos.com.br/dload/src/server"
)

type application struct {
	Server *server.Server
}

func InitializeApplication(properties config.Properties) (application, error) {
	handler := provideRouter()
	server := provideServer(handler, properties)
	mainApplication := application{Server: server}

	config.ConnectPostgres(properties.DB)
	repository.NewRepository()
	migrate.Migrate()

	return mainApplication, nil
}
