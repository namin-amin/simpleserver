package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

func (c *Config) GetEnvVarWithDefault(name string, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		value = c.GetString(name)
		if value == "" {
			return defaultValue
		}
	}
	return value
}

func NewConfig() *Config {
	c := &Config{
		Viper: viper.New(),
	}
	initialiseDefaultConfig(c)
	return c
}

func initialiseDefaultConfig(config *Config) {
	defaultConfigList := [3]string{"config.yaml", "config-dev.yaml", ".env"}

	for _, configs := range defaultConfigList {
		if fileExits(configs) {
			loadAndMergeConfigs(configs, config)
		}
	}
}

func loadAndMergeConfigs(in string, config *Config) {
	config.SetConfigFile(in)
	err := config.MergeInConfig()

	if err != nil {
		panic("could not load config \n" + err.Error())
	}
}

func fileExits(filepath string) bool {
	_, err := os.Stat(filepath)
	return !errors.Is(err, os.ErrNotExist)
}
