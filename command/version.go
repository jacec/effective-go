package command

import (
	"bytes"
	"fmt"

	"github.com/mitchellh/cli"
)

// VersionCommand is a Command implementation prints the version.
type VersionCommand struct {
	Revision          string
	Version           string
	VersionPrerelease string
	UI                cli.Ui
}

//Help returns the help for the command
func (c *VersionCommand) Help() string {
	return ""
}

//Run runs the command
func (c *VersionCommand) Run(_ []string) int {
	var versionString bytes.Buffer
	fmt.Fprintf(&versionString, "effective-go %s", c.Version)
	if c.VersionPrerelease != "" {
		fmt.Fprintf(&versionString, ".%s", c.VersionPrerelease)

		if c.Revision != "" {
			fmt.Fprintf(&versionString, " (%s)", c.Revision)
		}
	}

	c.UI.Output(versionString.String())
	return 0
}

//Synopsis resturns the synopsis for the command
func (c *VersionCommand) Synopsis() string {
	return "Prints the [app name] version"
}
