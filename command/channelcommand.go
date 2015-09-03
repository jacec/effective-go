package command

import (
	"flag"
	"strings"

	"github.com/mitchellh/cli"
)

// ChannelCommand is a Command implementation for specifying the command...
type ChannelCommand struct {
	UI cli.Ui
}

//Help returns the help for the command
func (c *ChannelCommand) Help() string {
	helpText := `
Usage: effective-go channels [options]

  A description of the channels is available here
  https://golang.org/doc/effective_go.html#concurrency

Options:
  -code-snipet = [1..]
`
	return strings.TrimSpace(helpText)
}

//Run runs the command
func (c *ChannelCommand) Run(args []string) int {

	var hclFile string
	cmdFlags := flag.NewFlagSet("channelcommand", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.UI.Output(c.Help()) }
	cmdFlags.StringVar(&hclFile, "code-snipet", "", "")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	c.UI.Output("ChannelCommand Complete")
	return 0
}

//Synopsis resturns the synopsis for the command
func (c *ChannelCommand) Synopsis() string {
	return "Run code snipets and get link for channels in effective go."
}
