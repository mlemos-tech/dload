package config

import (
	"flag"

	"github.com/joho/godotenv"
)

func UpEnvironmentConfig() (Properties, error) {
	var envfile string
	flag.StringVar(&envfile, "env-file", "../resource/dev.env", "Read in file of enviroment variables")
	flag.Parse()

	godotenv.Load(envfile)
	return ConfigEnvironment()
}
