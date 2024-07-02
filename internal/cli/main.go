package cli

import (
	"flag"
)

type Cli struct {
	ConfigFile string
}

func NewCli() (c *Cli, err error) {
	configFilePtr := flag.String("config", "example.yml", "config file, example: example.yml")

	flag.Parse()

	configFile := *configFilePtr

	c = &Cli{
		ConfigFile: configFile,
	}

	return
}
