package cmd

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type cmdSSHGet struct{}

func (cmd *cmdSSHGet) run(c *kingpin.ParseContext) error {
	return nil
}

// SSHGet declares the "ssh-get" sub command.
func SSHGet(app *kingpin.Application) {
	c := new(cmdSSHGet)

	app.Command("ssh-get", "Returns the SSH secret configuration").Action(c.run)
}