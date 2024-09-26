package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPServer      HTTPServer     `yaml:"http_server"`
	Swag            Swag           `yaml:"swag"`
	PostgresConfig  PostgresConfig `yaml:"postgres"`
	GracefulTimeout int            `yaml:"graceful_shutdown_timeout"`
	Env             string         `yaml:"env"`
	Logger          Logger         `yaml:"logger"`
	Cronjob         Cronjob        `yaml:"cronjob"`
}

type Logger struct {
	Path string `yaml:"path"`
}

func MustLoad() *Config {
	configPath := flag.String("config", "", "path to yaml configure file.")

	flag.Parse()

	if *configPath == "" {
		log.Fatalf("config command flag is not set")
	}

	if _, err := os.Stat(*configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	config := new(Config)

	if err := cleanenv.ReadConfig(*configPath, config); err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	return config
}
