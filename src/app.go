package main

import (
	"github.com/sirupsen/logrus"
	"mikaellemos.com.br/dload/src/config"
)

func main() {
	properties, err := config.UpEnvironmentConfig()

	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: invalid configuration")
	}

	config.AppConfigLog(properties)
	app, err := InitializeApplication(properties)

	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: cannot initialize server")
	}

	logrus.WithFields(logrus.Fields{"Host": properties.Server.Host, "Port": properties.Server.Port}).Info("starting the http server")
	app.Server.ListenAndServe()
}
