package config

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

// ProjectConfig is project config, it is should only include table
type ProjectConfig struct {
	LogConfig    LoggerConfig     `toml:"log"`
	ServerConfig ServerInfoConfig `toml:"server"`
}

// LoggerConfig include project logger related configuration
type LoggerConfig struct {
	LogPath           string `toml:"path"`
	LogLevel          string `toml:"level"`
	LogFileMaxSize    int    `toml:"size"`
	LogFileMaxAge     int    `toml:"age"`
	LogFileMaxBackups int    `toml:"backups"`
	LocalTime         bool   `toml:"local"`
	Compress          bool   `toml:"compress"`
}

type ServerInfoConfig struct {
	Host               string `toml:"host"`
	Port               string `toml:"port"`
	ApplicationID      string `toml:"application_id"`
	ApplicationName    string `toml:"application_name"`
	InstanceID         string `toml:"instance_id"`
	InstanceName       string `toml:"instance_name"`
	ApplicationVersion string `toml:"application_version"`
}

// InitConfig load config file content and parse
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
