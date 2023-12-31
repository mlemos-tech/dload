package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type (
	Properties struct {
		Server  Server
		Logging Logging
		DB      DB
	}

	Logging struct {
		Debug  bool `envconfig:"LOGS_DEBUG"`
		Trace  bool `envconfig:"LOGS_TRACE"`
		Color  bool `envconfig:"LOGS_COLOR"`
		Pretty bool `envconfig:"LOGS_PRETTY"`
		Text   bool `envconfig:"LOGS_TEXT"`
	}

	Server struct {
		Host string `envconfig:"SERVER_HOST" default:"localhost"`
		Port int    `envconfig:"SERVER_PORT" default:"8080"`
	}

	DB struct {
		Uri      string `envconfig:"DB_URI"`
		Database string `envconfig:"DB_NAME" default:"dload"`
	}
)

func ConfigEnvironment() (Properties, error) {
	cfg := Properties{}
	err := envconfig.Process("", &cfg)

	return cfg, err
}

func (c *Properties) String() string {
	out, _ := yaml.Marshal(c)
	return string(out)
}
