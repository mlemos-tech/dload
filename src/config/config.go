package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Server  Server
		Logging Logging
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
)

func ConfigEnvironment() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)

	return cfg, err
}

func (c *Config) String() string {
	out, _ := yaml.Marshal(c)
	return string(out)
}
