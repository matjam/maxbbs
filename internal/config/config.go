package config

import (
	"log/slog"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server struct {
		Sysop string
		Name  string

		Telnet struct {
			Host string
			Port int
		}
	}
}

const configPath = "cfg/max.toml"

var config *Config

func Get() *Config {
	if config == nil {
		config = &Config{}
		_, err := toml.DecodeFile(configPath, config)
		if err != nil {
			slog.Error("could not load config", "file", configPath, "error", err.Error())
			os.Exit(-1)
		}
	}

	return config
}
