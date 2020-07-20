package config

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

type ProjectConfig struct {
	LogConfig LoggerConfig `toml:"log"`
}

type LoggerConfig struct {
	LogPath           string `toml:"path"`
	LogLevel          string `toml:"level"`
	LogFileMaxSize    int    `toml:"size"`
	LogFileMaxAge     int    `toml:"age"`
	LogFileMaxBackups int    `toml:"backups"`
	LocalTime         bool   `toml:"local"`
	Compress          bool   `toml:"compress"`
}

func InitConfig(configFile string) (ProjectConfig, error) {
	projectConfigContent := ProjectConfig{}

	configFileContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		return projectConfigContent, err
	}

	err = toml.Unmarshal(configFileContent, &projectConfigContent)
	if err != nil {
		return projectConfigContent, err
	}

	return projectConfigContent, nil
}
