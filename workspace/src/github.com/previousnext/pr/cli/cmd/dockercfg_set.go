package cmd

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type cmdDockerCfgSet struct{}

func (cmd *cmdDockerCfgSet) run(c *kingpin.ParseContext) error {
	return nil
}

// DockerCfgSet declares the "dockercfg-set" sub command.
func DockerCfgSet(app *kingpin.Application) {
	c := new(cmdDockerCfgSet)

	app.Command("dockercfg-set", "Updates the Docker secret configuration").Action(c.run)
}
