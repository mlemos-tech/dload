package main

import (
	"mikaellemos.com.br/dload/src/config"
	"mikaellemos.com.br/dload/src/migrate"
	"mikaellemos.com.br/dload/src/server"
)

type application struct {
	Server *server.Server
}

func InitializeApplication(propertie config.Properties) (application, error) {
	handler := provideRouter()
	server := provideServer(handler, propertie)
	mainApplication := application{Server: server}

	config.ConnectPostgres("host=localhost user=dload password=123 dbname=dload port=5432")
	migrate.Migrate()

	return mainApplication, nil
}
