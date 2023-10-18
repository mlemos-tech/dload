package config

import (
	"github.com/sirupsen/logrus"
)

func AppConfigLog(c Properties) {
	if c.Logging.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if c.Logging.Trace {
		logrus.SetLevel(logrus.TraceLevel)
	}

	if c.Logging.Text {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   c.Logging.Color,
			DisableColors: !c.Logging.Color,
		})

	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: c.Logging.Pretty,
		})
	}
}
