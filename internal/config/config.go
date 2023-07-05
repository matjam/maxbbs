package config

import (
	"os"

	"github.com/goccy/go-yaml"
	"github.com/gookit/slog"
)

type Config struct {
	SysopName string `yaml:"sysop_name"`
	SysName   string `yaml:"sys_name"`
}

const configPath = "cfg/max.yaml"

var config *Config

func Get() *Config {
	if config == nil {
		config = &Config{}
		data, err := os.ReadFile(configPath)
		if err != nil {
			slog.Panicf("could not load config at %v: %v", configPath, err.Error())
		}
		if err := yaml.Unmarshal(data, config); err != nil {
			slog.Panicf("could not parse config at %v: %v", configPath, err.Error())
		}
	}

	return config
}
