package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

// SubCommand is a Command implementation for specifying the command...
type SubCommand struct {
	UI cli.Ui
}

//Help returns the help for the command
func (c *SubCommand) Help() string {
	helpText := `
Usage: [app name] [subcommand] [options]

  A description of the [subcommand]

Options:

  -key=val                     description of the key
`
	return strings.TrimSpace(helpText)
}

//Run runs the command
func (c *SubCommand) Run(args []string) int {

	var hclFile string
	cmdFlags := flag.NewFlagSet("subcommand", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.UI.Output(c.Help()) }
	cmdFlags.StringVar(&hclFile, "key", "", "")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	c.UI.Output("Subcommand Complete")
	return 0
}

//Synopsis resturns the synopsis for the command
func (c *SubCommand) Synopsis() string {
	return "Synopsis of [subcommand]"
}
