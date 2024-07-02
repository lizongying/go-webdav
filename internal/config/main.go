package config

import (
	"errors"
	"github.com/lizongying/go-webdav/internal/cli"
	"gopkg.in/yaml.v3"
	"os"
)

type Server struct {
	Host string `yaml:"host" json:"host"`
	Cert string `yaml:"cert" json:"cert"`
	Key  string `yaml:"key" json:"key"`
}

type Config struct {
	Dirs   []string `yaml:"dirs" json:"dirs"`
	Server Server   `yaml:"server" json:"server"`
}

func (c *Config) GetDirs() []string {
	return c.Dirs
}

func (c *Config) GetServer() Server {
	return c.Server
}

func (c *Config) LoadConfig(configPath string) (err error) {
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(configData, c)
	return
}

func NewConfig(cli *cli.Cli) (config *Config, err error) {
	config = new(Config)
	configFile := cli.ConfigFile
	if configFile == "" {
		err = errors.New("config file is empty")
		return
	}

	err = config.LoadConfig(configFile)
	return
}
