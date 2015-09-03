package command

import (
	"github.com/mitchellh/cli"
	"testing"
)

func TestVersionCommand_implements(t *testing.T) {
	var _ cli.Command = &VersionCommand{}
}

func TestVersionCommandRun(t *testing.T) {

	ui := new(cli.MockUi)
	c := &VersionCommand{UI: ui}
	args := []string{}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

}
